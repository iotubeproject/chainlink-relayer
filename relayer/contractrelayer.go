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
	"github.com/go-lark/lark"
	"github.com/pkg/errors"

	"github.com/iotexproject/chainlink-relayer/contract"
)

const (
	FetchData = iota + 1
	FetchConfig
)

type (
	AggregatorConfig struct {
		AggregatorAddr string `json:"aggregatorAddr" yaml:"aggregatorAddr"`
		IsDummy        bool   `json:"isDummy" yaml:"isDummy"`
	}
	pair struct {
		sourceAggregatorAddr common.Address
		shadowAggregatorAddr common.Address
		shadowAggregator     interface{}
		dummy                bool
	}
	contractRelayer struct {
		abstractRelayer
		mode                   int
		balanceLowerbound      *big.Int
		lastProcessBlockHeight uint64
		batchSize              uint64
		sourceClient           *ethclient.Client
		recorder               *Recorder
		pairs                  []pair
		isDummyAggregator      map[common.Address]bool
	}
)

func NewContractRelayer(
	privateKey string,
	mode int,
	startHeight uint64,
	recorder *Recorder,
	aggregators map[string]AggregatorConfig,
	sourceClient *ethclient.Client,
	targetChainID uint32,
	targetClient *ethclient.Client,
	hookUrl string,
	batchSize uint64,
) (Relayer, error) {
	pairs := []pair{}
	isDummyAggregator := make(map[common.Address]bool)
	for aggregatorAddr, config := range aggregators {
		fmt.Println(aggregatorAddr, "=>", config.AggregatorAddr, config.IsDummy)
		if config.IsDummy {
			dummyAggragator, err := contract.NewDummyAggregator(common.HexToAddress(config.AggregatorAddr), targetClient)
			if err != nil {
				return nil, err
			}
			pairs = append(pairs, pair{
				sourceAggregatorAddr: common.HexToAddress(aggregatorAddr),
				shadowAggregatorAddr: common.HexToAddress(config.AggregatorAddr),
				shadowAggregator:     dummyAggragator,
				dummy:                true,
			})
			isDummyAggregator[common.HexToAddress(aggregatorAddr)] = true
		} else {
			shadowAggregator, err := contract.NewShadowAggregator(common.HexToAddress(config.AggregatorAddr), targetClient)
			if err != nil {
				return nil, err
			}
			pairs = append(pairs, pair{
				sourceAggregatorAddr: common.HexToAddress(aggregatorAddr),
				shadowAggregatorAddr: common.HexToAddress(config.AggregatorAddr),
				shadowAggregator:     shadowAggregator,
				dummy:                false,
			})
			isDummyAggregator[common.HexToAddress(aggregatorAddr)] = false
		}
	}
	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}
	var notificationBot *lark.Bot
	if hookUrl != "" {
		notificationBot = lark.NewNotificationBot(hookUrl)
	}
	lowerbound, _ := new(big.Int).SetString("10000000000000000000", 10)

	return &contractRelayer{
		balanceLowerbound: lowerbound,
		mode:              mode,
		abstractRelayer: abstractRelayer{
			notificationBot:    notificationBot,
			targetChainID:      big.NewInt(int64(targetChainID)),
			gasLimit:           5000000,
			gasPriceUpperBound: big.NewInt(1000000000000000000),
			privateKey:         pk,
			recorder:           recorder,
			targetClient:       targetClient,
		},
		pairs:                  pairs,
		isDummyAggregator:      isDummyAggregator,
		lastProcessBlockHeight: startHeight,
		batchSize:              batchSize,
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

func (relayer *contractRelayer) pullData(
	ctx context.Context,
	tipHeight uint64,
	sourceAggregatorAddr, shadowAggregatorAddr common.Address,
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
	endHeight := startHeight + relayer.batchSize
	if tipHeight < endHeight {
		endHeight = tipHeight
	}
	fmt.Println("Fetching data from", startHeight, "to", endHeight)
	logs, err := relayer.sourceClient.FilterLogs(ctx,
		ethereum.FilterQuery{
			FromBlock: new(big.Int).SetUint64(startHeight),
			ToBlock:   new(big.Int).SetUint64(endHeight),
			Addresses: []common.Address{sourceAggregatorAddr},
			Topics: [][]common.Hash{{
				ocr2aggregatorABI.Events[EventNewTransmission].ID,
				// ocr2aggregatorABI.Events[EventConfigSet].ID,
				aggregatorABI.Events[EventNewRound].ID,
				aggregatorABI.Events[EventConfigSet].ID}},
		},
	)
	fmt.Println("Fetched", len(logs), "logs", err)
	if err != nil {
		return err
	}
	transmissions := []*Transmission{}
	rounds := []*Round{}
	configs := []*ConfigSet{}
	for _, log := range logs {
		fmt.Println(log.BlockNumber, log.TxHash.Hex(), log.Topics[0].Hex())
		switch log.Topics[0] {
		case ocr2aggregatorABI.Events[EventNewTransmission].ID:
			if !relayer.isDummyAggregator[log.Address] {
				continue
			}
			header, err := relayer.sourceClient.HeaderByNumber(ctx, new(big.Int).SetUint64(log.BlockNumber))
			if err != nil {
				return err
			}
			transmission, err := NewTransmission(log, header.Time)
			if err != nil {
				return err
			}
			transmissions = append(transmissions, transmission)
		case aggregatorABI.Events[EventNewRound].ID:
			if relayer.isDummyAggregator[log.Address] {
				continue
			}
			tx, _, err := relayer.sourceClient.TransactionByHash(ctx, log.TxHash)
			if err != nil {
				return err
			}
			round, err := NewRound(log, tx)
			if err != nil {
				return err
			}
			rounds = append(rounds, round)
		case aggregatorABI.Events[EventConfigSet].ID:
			config, err := NewConfigSet(log)
			if err != nil {
				return err
			}
			configs = append(configs, config)
		}
	}
	if len(rounds) > 0 || len(configs) > 0 {
		fmt.Printf("%s: %d rounds and %d configs fetched from %d to %d\n", sourceAggregatorAddr, len(rounds), len(configs), startHeight, endHeight)
	}

	return relayer.recorder.PutRounds(shadowAggregatorAddr.String(), strconv.FormatUint(endHeight, 10), transmissions, rounds, configs)
}

func (relayer *contractRelayer) Produce(ctx context.Context) error {
	tipHeight, err := relayer.tipHeight(ctx)
	if err != nil {
		return err
	}
	for _, p := range relayer.pairs {
		if err := relayer.pullData(ctx, tipHeight, p.sourceAggregatorAddr, p.shadowAggregatorAddr); err != nil {
			fmt.Printf("failed to pull data for %s: %+v\n", p.shadowAggregatorAddr, err)
			return err
		}
	}
	return nil
}

func (relayer *contractRelayer) consume(
	ctx context.Context,
	sourceAggregatorAddr common.Address,
	shadowAggregator interface{},
	isDummy bool,
) (bool, error) {
	if isDummy {
		return relayer.consumeDummy(ctx, sourceAggregatorAddr, shadowAggregator.(*contract.DummyAggregator))
	}
	return relayer.consumeShadow(ctx, sourceAggregatorAddr, shadowAggregator.(*contract.ShadowAggregator))
}

func (relayer *contractRelayer) consumeDummy(ctx context.Context, sourceAggregatorAddr common.Address, dummyAggregator *contract.DummyAggregator) (bool, error) {
	toConfirm, err := relayer.recorder.TransmissionToConfirm(sourceAggregatorAddr.String())
	if err != nil {
		return false, err
	}
	if toConfirm != nil {
		nonce, err := relayer.nonceAt(ctx)
		if err != nil {
			return false, err
		}
		switch receipt, err := relayer.targetClient.TransactionReceipt(ctx, common.HexToHash(toConfirm.RelayTxHash)); errors.Cause(err) {
		case nil:
			if receipt == nil {
				if toConfirm.Nonce <= nonce && toConfirm.CreatedAt.Add(10*time.Minute).Before(time.Now()) {
					return false, relayer.recorder.ResetTransmission(toConfirm.ID)
				}
				return false, nil
			}
			if receipt.Status != 1 {
				if err := relayer.recorder.FailTransmission(toConfirm.ID); err != nil {
					return false, err
				}
			} else {
				if err := relayer.recorder.ConfirmTransmission(toConfirm.ID); err != nil {
					return false, err
				}
			}
		case ethereum.NotFound:
			if toConfirm.Nonce <= nonce && toConfirm.CreatedAt.Add(10*time.Minute).Before(time.Now()) {
				return false, relayer.recorder.ResetTransmission(toConfirm.ID)
			}
			return false, nil
		default:
			return false, err
		}
	}
	var toRelay *Transmission
	if (relayer.mode & FetchData) != 0 {
		toRelay, err = relayer.recorder.TransmissionToRelay(sourceAggregatorAddr.String())
		if err != nil {
			return false, err
		}
	}
	var configToRelay *ConfigSet
	if (relayer.mode & FetchConfig) != 0 {
		configToRelay, err = relayer.recorder.ConfigToRelay(sourceAggregatorAddr.String())
		if err != nil {
			return false, err
		}
	}
	if toRelay == nil && configToRelay == nil {
		return false, nil
	}
	if err := relayer.checkBalance(ctx); err != nil {
		return false, err
	}
	return true, relayer.relayDummy(ctx, dummyAggregator, toRelay, configToRelay)
}

func (relayer *contractRelayer) consumeShadow(ctx context.Context, sourceAggregatorAddr common.Address, shadowAggregator *contract.ShadowAggregator) (bool, error) {
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
	// TODO: confirm ConfigSet
	var roundToRelay *Round
	if (relayer.mode & FetchData) != 0 {
		roundToRelay, err = relayer.recorder.RoundsToRelay(sourceAggregatorAddr.String())
		if err != nil {
			return false, err
		}
	}
	var configToRelay *ConfigSet
	if (relayer.mode & FetchConfig) != 0 {
		configToRelay, err = relayer.recorder.ConfigToRelay(sourceAggregatorAddr.String())
		if err != nil {
			return false, err
		}
	}
	if roundToRelay == nil && configToRelay == nil {
		return false, nil
	}
	if err := relayer.checkBalance(ctx); err != nil {
		return false, err
	}

	return true, relayer.relayShadow(ctx, shadowAggregator, roundToRelay, configToRelay)
}

func (relayer *contractRelayer) Consume(ctx context.Context) error {
	for _, p := range relayer.pairs {
		relayed, err := relayer.consume(ctx, p.sourceAggregatorAddr, p.shadowAggregator, p.dummy)
		if err != nil {
			return err
		}
		if relayed {
			time.Sleep(10 * time.Second)
		}
	}

	return nil
}

func (relayer *contractRelayer) checkBalance(ctx context.Context) error {
	balance, err := relayer.balance(ctx)
	if err != nil {
		return err
	}
	if balance.Cmp(relayer.balanceLowerbound) < 0 {
		relayer.alert(fmt.Sprintf(
			"Running out of iotx, balance %d is lower than %d, please refill %s right away",
			balance,
			relayer.balanceLowerbound,
			relayer.address(),
		))
	}
	return nil
}

func (relayer *contractRelayer) relayDummy(
	ctx context.Context,
	dummyAggregator *contract.DummyAggregator,
	round *Transmission,
	config *ConfigSet,
) error {
	owner, err := dummyAggregator.Owner(nil)
	if err != nil {
		return err
	}
	if relayer.address() != owner {
		fmt.Printf("relayer %s is not owner %s\n", relayer.address(), owner)
		return nil
	}
	opts, err := relayer.transactionOpts(ctx)
	if err != nil {
		return err
	}
	if round != nil {
		latest, err := dummyAggregator.LatestRoundData(nil)
		if err != nil {
			return err
		}
		if latest.RoundId.Uint64() >= uint64(round.AggregatorRoundId) {
			fmt.Printf("Skip dummy <%s %d> which is lower than round %d on chain\n", round.Aggregator, round.AggregatorRoundId, latest.RoundId)
			relayer.recorder.SkipRound(round.ID)
			return nil
		}
		fmt.Printf("Submitting dummy <%s, %d>...\n", round.Aggregator, round.AggregatorRoundId)
		_, answer, _, _, _, epochAndRound, _, err := round.FormatData()
		if err != nil {
			return err
		}
		tx, err := dummyAggregator.AddTransmission(
			opts,
			epochAndRound,
			round.AggregatorRoundId,
			answer,
			round.ObservationsTimestamp,
			round.TransmissionTimestamp,
		)
		if err != nil {
			relayer.alert(fmt.Sprintf("failed to relay dummy <%s, %d>, %+v\n", round.Aggregator, round.AggregatorRoundId, err))
			return err
		}
		return relayer.recorder.SetTransmissionRelayTxHash(round.ID, tx.Hash(), relayer.address(), tx.Nonce())
	}
	if config != nil {
		fmt.Printf("Submitting dummy config <%s, %d>...\n", config.Aggregator, config.ConfigCount)
		tx, err := dummyAggregator.ResetEpochAndRound(
			opts,
		)
		if err != nil {
			relayer.alert(fmt.Sprintf("failed to relay dummy config <%s, %d>, %+v\n", config.Aggregator, config.ConfigCount, err))
			return err
		}
		return relayer.recorder.SetConfigRelayTxHash(config.ID, tx.Hash(), opts.From, tx.Nonce())
	}
	return nil
}

func (relayer *contractRelayer) relayShadow(
	ctx context.Context,
	shadowAggregator *contract.ShadowAggregator,
	round *Round,
	config *ConfigSet,
) error {
	digest, err := shadowAggregator.ConfigDigest(nil)
	if err != nil {
		return err
	}
	opts, err := relayer.transactionOpts(ctx)
	if err != nil {
		return err
	}
	if round != nil {
		report, rs, ss, vs, err := round.FormatData()
		if err != nil {
			return err
		}
		if bytes.Compare(report[11:27], digest[:]) == 0 {
			roundId, err := shadowAggregator.LatestRoundId(nil)
			if err != nil {
				err = errors.Wrapf(err, "failed to get round id from shadow aggregator")
				relayer.alert(err.Error())
				return err
			}
			if round.Number <= roundId.Uint64() {
				fmt.Printf("Skip <%s %d> which is lower than round %d on chain\n", round.Aggregator, round.Number, roundId)
				relayer.recorder.SkipRound(round.ID)
				return nil
			}
			fmt.Printf("Submitting <%s, %d> %d...\n", round.Aggregator, round.Number, roundId)
			tx, err := shadowAggregator.Submit(
				opts,
				report,
				rs,
				ss,
				vs,
			)
			if err != nil {
				relayer.alert(fmt.Sprintf("failed to relay <%s, %d>, %+v\n", round.Aggregator, round.Number, err))
				return err
			}
			return relayer.recorder.SetRoundRelayTxHash(round.ID, tx.Hash(), opts.From, tx.Nonce())
		}
		if config == nil {
			notification := fmt.Sprintf("Digest inconsistency %x vs %x for <%s, %d>", report[11:27], digest, round.Aggregator, round.Number)
			fmt.Println(notification)
			if round.CreatedAt.Add(time.Minute).Before(time.Now()) {
				relayer.alert(notification)
			}
			return nil
		}
	}
	owner, err := shadowAggregator.Owner(nil)
	if err != nil {
		return err
	}
	if relayer.address() != owner {
		fmt.Printf("relayer %s is not owner %s, waiting for owner to set config\n", relayer.address(), owner)
		return nil
	}
	signers, transmitters, threshold, version, encoded, err := config.FormatData()
	if err != nil {
		return err
	}
	tx, err := shadowAggregator.SetConfig(opts, signers, transmitters, threshold, version, encoded)
	if err != nil {
		relayer.alert(fmt.Sprintf("failed to relay config <%s, %d>, %+v\n", config.Aggregator, config.ConfigCount, err))
		return err
	}
	return relayer.recorder.SetConfigRelayTxHash(config.ID, tx.Hash(), opts.From, tx.Nonce())
}
