package relayer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/iotexproject/chainlink-relayer/contract"
	"github.com/pkg/errors"
)

type (
	Relayer struct {
		targetChainID          *big.Int
		gasPriceUpperBound     *big.Int
		lastProcessBlockHeight uint64
		sourceClient           *ethclient.Client
		targetClient           *ethclient.Client
		recorder               *Recorder
		shadowAggregators      map[common.Address]*contract.ShadowAggregator
		query                  ethereum.FilterQuery
		privateKey             *ecdsa.PrivateKey
		gasLimit               uint64
	}
)

func NewRelayer(
	privateKey string,
	startHeight uint64,
	recorder *Recorder,
	pairs map[common.Address]common.Address,
	sourceClient *ethclient.Client,
	targetClient *ethclient.Client,
) (*Relayer, error) {
	shadowAggregators := map[common.Address]*contract.ShadowAggregator{}
	aggregatorAddrs := []common.Address{}
	for aggregatorAddr, shadowAggregatorAddr := range pairs {
		shadowAggregator, err := contract.NewShadowAggregator(shadowAggregatorAddr, targetClient)
		if err != nil {
			return nil, err
		}
		shadowAggregators[aggregatorAddr] = shadowAggregator
		aggregatorAddrs = append(aggregatorAddrs, aggregatorAddr)
	}
	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}

	return &Relayer{
		targetChainID:          big.NewInt(4690),
		gasLimit:               1000000,
		gasPriceUpperBound:     big.NewInt(1000000000000000000),
		privateKey:             pk,
		shadowAggregators:      shadowAggregators,
		lastProcessBlockHeight: startHeight,
		recorder:               recorder,
		sourceClient:           sourceClient,
		targetClient:           targetClient,
		query: ethereum.FilterQuery{
			Addresses: aggregatorAddrs,
			Topics:    [][]common.Hash{{aggregatorABI.Events[EventNewRound].ID}},
		},
	}, nil
}

func (relayer *Relayer) Produce(ctx context.Context) error {
	query := relayer.query
	startHeight := relayer.lastProcessBlockHeight
	syncedHeight, err := relayer.recorder.SyncedHeight()
	if err != nil {
		return err
	}
	if syncedHeight > startHeight {
		startHeight = syncedHeight
	}
	startHeight += 1
	tipHeight, err := relayer.sourceClient.BlockNumber(ctx)
	if err != nil {
		return err
	}
	if tipHeight < 20 {
		return nil
	}
	tipHeight -= 20
	if tipHeight < startHeight {
		return nil
	}
	endHeight := startHeight + 99
	if tipHeight < endHeight {
		endHeight = tipHeight
	}
	query.FromBlock = new(big.Int).SetUint64(startHeight)
	query.ToBlock = new(big.Int).SetUint64(endHeight)
	logs, err := relayer.sourceClient.FilterLogs(ctx, query)
	if err != nil {
		return err
	}
	rounds := make([]*Round, len(logs))
	for i, log := range logs {
		tx, _, err := relayer.sourceClient.TransactionByHash(ctx, log.TxHash)
		if err != nil {
			return err
		}
		round, err := NewRound(log, tx)
		if err != nil {
			return err
		}
		rounds[i] = round
	}
	fmt.Printf("%d fetched from %d to %d\n", len(rounds), query.FromBlock, query.ToBlock)
	return relayer.recorder.Put(query.ToBlock.Uint64(), rounds)
}

func (relayer *Relayer) Consume(ctx context.Context) error {
	aggregators := []string{}
	for aggregator := range relayer.shadowAggregators {
		aggregators = append(aggregators, aggregator.String())
	}
	roundsToConfirm, err := relayer.recorder.RoundsToConfirm(aggregators...)
	if err != nil {
		return err
	}
	roundsToRelay, err := relayer.recorder.RoundsToRelay(aggregators...)
	if err != nil {
		return err
	}
	fmt.Printf("%d to confirm, %d to relay\n", len(roundsToConfirm), len(roundsToRelay))
	notConfirmed := map[string]bool{}
	for _, round := range roundsToConfirm {
		notConfirmed[round.Aggregator] = true
		_, err := relayer.targetClient.TransactionReceipt(ctx, common.HexToHash(round.RelayTxHash))
		switch errors.Cause(err) {
		case nil:
			if err := relayer.recorder.Confirm(round.ID); err != nil {
				fmt.Printf("failed to confirm %+v, %+v\n", round, err)
			}
			notConfirmed[round.Aggregator] = false
		case ethereum.NotFound:
			continue
		default:
			fmt.Printf("failed fetch receipt for %s, %+v\n", round.RelayTxHash, err)
		}
	}
	for _, round := range roundsToRelay {
		if notConfirmed[round.Aggregator] {
			fmt.Printf("skip relay <%s, %d>, waiting to confirm previous round\n", round.Aggregator, round.Number)
			continue
		}
		notConfirmed[round.Aggregator] = true
		opts, err := relayer.transactionOpts(ctx)
		if err != nil {
			return err
		}
		report, rs, ss, vs, err := round.FormatData()
		if err != nil {
			return err
		}
		fmt.Printf("Submitting <%s, %d>...\n", round.Aggregator, round.Number)
		tx, err := relayer.shadowAggregators[common.HexToAddress(round.Aggregator)].Submit(
			opts,
			report,
			rs,
			ss,
			vs,
		)
		if err != nil {
			fmt.Printf("failed to relay <%s, %d>, %+v\n", round.Aggregator, round.Number, err)
			continue
		}
		if err := relayer.recorder.SetRelayTxHash(round.ID, tx.Hash()); err != nil {
			return err
		}
	}

	return nil
}

func (relayer *Relayer) transactionOpts(ctx context.Context) (*bind.TransactOpts, error) {
	relayerAddr := crypto.PubkeyToAddress(relayer.privateKey.PublicKey)

	opts, err := bind.NewKeyedTransactorWithChainID(relayer.privateKey, relayer.targetChainID)
	if err != nil {
		return nil, err
	}
	opts.Value = big.NewInt(0)
	opts.GasLimit = relayer.gasLimit
	gasPrice, err := relayer.targetClient.SuggestGasPrice(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get suggested gas price")
	}
	if gasPrice.Cmp(relayer.gasPriceUpperBound) >= 0 {
		return nil, errors.Wrapf(ErrGasPriceTooHigh, "suggested gas price %d > limit %d", gasPrice, relayer.gasPriceUpperBound)
	}
	opts.GasPrice = gasPrice
	balance, err := relayer.targetClient.BalanceAt(ctx, relayerAddr, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get balance of operator account")
	}
	gasFee := new(big.Int).Mul(new(big.Int).SetUint64(opts.GasLimit), opts.GasPrice)
	if gasFee.Cmp(balance) > 0 {
		return nil, errors.Errorf("insuffient balance for gas fee")
	}
	nonce, err := relayer.targetClient.PendingNonceAt(ctx, relayerAddr)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to fetch pending nonce for %s", relayerAddr)
	}
	opts.Nonce = new(big.Int).SetUint64(nonce)

	return opts, nil
}
