package relayer

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/iotexproject/chainlink-relayer/contract"
)

type (
	pair struct {
		sourceAggregatorAddr common.Address
		shadowAggregatorAddr common.Address
		shadowAggregator     *contract.ShadowAggregator
	}
	contractRelayer struct {
		abstractRelayer
		lastProcessBlockHeight uint64
		sourceClient           *ethclient.Client
		recorder               *Recorder
		pairs                  []pair
	}
)

func NewContractRelayer(
	privateKey string,
	startHeight uint64,
	recorder *Recorder,
	aggregators map[common.Address]common.Address,
	sourceClient *ethclient.Client,
	targetChainID uint32,
	targetClient *ethclient.Client,
) (Relayer, error) {
	pairs := []pair{}
	for aggregatorAddr, shadowAggregatorAddr := range aggregators {
		shadowAggregator, err := contract.NewShadowAggregator(shadowAggregatorAddr, targetClient)
		if err != nil {
			return nil, err
		}
		pairs = append(pairs, pair{
			sourceAggregatorAddr: aggregatorAddr,
			shadowAggregatorAddr: shadowAggregatorAddr,
			shadowAggregator:     shadowAggregator,
		})
	}
	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}

	return &contractRelayer{
		abstractRelayer: abstractRelayer{
			targetChainID:      big.NewInt(int64(targetChainID)),
			gasLimit:           1000000,
			gasPriceUpperBound: big.NewInt(1000000000000000000),
			privateKey:         pk,
			recorder:           recorder,
			targetClient:       targetClient,
		},
		pairs:                  pairs,
		lastProcessBlockHeight: startHeight,
		recorder:               recorder,
		sourceClient:           sourceClient,
	}, nil
}

func (relayer *contractRelayer) tipHeight(ctx context.Context) (uint64, error) {
	tipHeight, err := relayer.sourceClient.BlockNumber(ctx)
	if err != nil {
		return 0, err
	}
	if tipHeight < 20 {
		return 0, errors.New("chain is not ready")
	}
	return tipHeight - 20, nil
}

func (relayer *contractRelayer) pullRounds(
	ctx context.Context,
	tipHeight uint64,
	sourceAggregatorAddr, shadowAggregatorAddr common.Address,
	shadowAggregator *contract.ShadowAggregator,
) error {
	value, err := relayer.recorder.Value(shadowAggregatorAddr.String())
	if err != nil {
		return err
	}
	syncedHeight := uint64(0)
	if value != "" {
		syncedHeight, err = strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
	}
	startHeight := relayer.lastProcessBlockHeight
	if syncedHeight > startHeight {
		startHeight = syncedHeight + 1
	}
	if tipHeight < startHeight {
		return nil
	}
	endHeight := startHeight + 99
	if tipHeight < endHeight {
		endHeight = tipHeight
	}
	logs, err := relayer.sourceClient.FilterLogs(ctx,
		ethereum.FilterQuery{
			FromBlock: new(big.Int).SetUint64(startHeight),
			ToBlock:   new(big.Int).SetUint64(endHeight),
			Addresses: []common.Address{sourceAggregatorAddr},
			Topics:    [][]common.Hash{{aggregatorABI.Events[EventNewRound].ID}},
		},
	)
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
	if len(rounds) > 0 {
		fmt.Printf("%s: %d fetched from %d to %d\n", sourceAggregatorAddr, len(rounds), startHeight, endHeight)
	}

	return relayer.recorder.PutRounds(shadowAggregatorAddr.String(), strconv.FormatUint(endHeight, 10), rounds)
}

func (relayer *contractRelayer) Produce(ctx context.Context) error {
	tipHeight, err := relayer.tipHeight(ctx)
	if err != nil {
		return err
	}
	for _, p := range relayer.pairs {
		if err := relayer.pullRounds(ctx, tipHeight, p.sourceAggregatorAddr, p.shadowAggregatorAddr, p.shadowAggregator); err != nil {
			return err
		}
	}
	return nil
}

func (relayer *contractRelayer) consume(
	ctx context.Context,
	sourceAggregatorAddr, shadowAggregatorAddr common.Address,
	shadowAggregator *contract.ShadowAggregator,
) (bool, error) {
	roundToConfirm, err := relayer.recorder.RoundToConfirm(sourceAggregatorAddr.String())
	if err != nil {
		return false, err
	}
	if roundToConfirm != nil {
		nonce, err := relayer.nonceAt(ctx)
		if err != nil {
			return false, err
		}
		// TODO: read contract, if already confirmed, skip
		switch receipt, err := relayer.targetClient.TransactionReceipt(ctx, common.HexToHash(roundToConfirm.RelayTxHash)); errors.Cause(err) {
		case nil:
			if receipt == nil {
				if roundToConfirm.Nonce <= nonce && roundToConfirm.CreatedAt.Add(10*time.Minute).Before(time.Now()) {
					return false, relayer.recorder.ResetRound(roundToConfirm.ID)
				}
				return false, nil
			}
			if receipt.Status != 1 {
				if err := relayer.recorder.FailRound(roundToConfirm.ID); err != nil {
					return false, err
				}
			} else {
				if err := relayer.recorder.ConfirmRound(roundToConfirm.ID); err != nil {
					return false, err
				}
			}
		case ethereum.NotFound:
			if roundToConfirm.Nonce <= nonce && roundToConfirm.CreatedAt.Add(10*time.Minute).Before(time.Now()) {
				return false, relayer.recorder.ResetRound(roundToConfirm.ID)
			}
			return false, nil
		default:
			return false, err
		}
	}
	roundToRelay, err := relayer.recorder.RoundsToRelay(sourceAggregatorAddr.String())
	if err != nil {
		return false, err
	}
	if roundToRelay == nil {
		return false, nil
	}

	return true, relayer.relay(ctx, shadowAggregator, roundToRelay)
}

func (relayer *contractRelayer) Consume(ctx context.Context) error {
	for _, p := range relayer.pairs {
		relayed, err := relayer.consume(ctx, p.sourceAggregatorAddr, p.shadowAggregatorAddr, p.shadowAggregator)
		if err != nil {
			return err
		}
		if relayed {
			time.Sleep(10 * time.Second)
		}
	}

	return nil
}

func (relayer *contractRelayer) relay(ctx context.Context, shadowAggregator *contract.ShadowAggregator, round *Round) error {
	opts, err := relayer.transactionOpts(ctx)
	if err != nil {
		return err
	}
	report, rs, ss, vs, err := round.FormatData()
	if err != nil {
		return err
	}
	digest, err := shadowAggregator.ConfigDigest(nil)
	if err != nil {
		return err
	}
	if bytes.Compare(report[11:27], digest[:]) != 0 {
		// TODO: alert admin
		fmt.Printf("Digest inconsistency %x vs %x for <%s, %d>\n", report[11:27], digest, round.Aggregator, round.Number)
		return nil
	}
	fmt.Printf("Submitting <%s, %d>...\n", round.Aggregator, round.Number)
	tx, err := shadowAggregator.Submit(
		opts,
		report,
		rs,
		ss,
		vs,
	)
	if err != nil {
		fmt.Printf("failed to relay <%s, %d>, %+v\n", round.Aggregator, round.Number, err)
		return err
	}
	return relayer.recorder.SetRoundRelayTxHash(round.ID, tx.Hash(), opts.From, tx.Nonce())
}
