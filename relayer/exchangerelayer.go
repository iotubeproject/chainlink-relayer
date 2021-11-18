package relayer

import (
	"context"
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/iotexproject/chainlink-relayer/contract"
	"github.com/iotexproject/chainlink-relayer/exchange"
)

type (
	exchangeRelayer struct {
		abstractRelayer
		exchanges      []exchange.Exchange
		aggregatorAddr common.Address
		aggregator     *contract.NaiveAggregator
	}
)

func NewExchangeRelayer(
	privateKey string,
	exchanges []exchange.Exchange,
	recorder *Recorder,
	aggregatorAddr common.Address,
	targetClient *ethclient.Client,
) (Relayer, error) {
	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}
	aggregator, err := contract.NewNaiveAggregator(aggregatorAddr, targetClient)
	if err != nil {
		return nil, err
	}
	return &exchangeRelayer{
		abstractRelayer: abstractRelayer{targetChainID: big.NewInt(4690),
			gasLimit:           1000000,
			gasPriceUpperBound: big.NewInt(1000000000000000000),
			privateKey:         pk,
			recorder:           recorder,
			targetClient:       targetClient,
		},
		exchanges:      exchanges,
		aggregatorAddr: aggregatorAddr,
		aggregator:     aggregator,
	}, nil
}

func (relayer *exchangeRelayer) Produce(context.Context) error {
	latestRoundId, err := relayer.aggregator.LatestRoundId(nil)
	if err != nil {
		return err
	}
	latestRoundData := struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	}{
		RoundId:         big.NewInt(0),
		Answer:          big.NewInt(0),
		StartedAt:       big.NewInt(0),
		UpdatedAt:       big.NewInt(0),
		AnsweredInRound: big.NewInt(0),
	}
	if latestRoundId.Cmp(big.NewInt(0)) != 0 {
		latestRoundData, err = relayer.aggregator.LatestRoundData(nil)
		if err != nil {
			return err
		}
	}
	number := latestRoundData.RoundId.Uint64()
	now := time.Now()
	latestRoundTime := time.Unix(latestRoundData.StartedAt.Int64(), 0)
	latestRoundAnswer := latestRoundData.Answer.Int64()
	var prices sort.Float64Slice
	for _, exchange := range relayer.exchanges {
		price, err := exchange.Price()
		if err != nil {
			return err
		}
		prices = append(prices, price)
	}
	prices.Sort()
	answer := int64(prices[len(prices)/2] * 1000000)
	diff := latestRoundAnswer - answer
	if diff < 0 {
		diff = -diff
	}
	if diff*200 < latestRoundAnswer && now.Before(latestRoundTime.Add(time.Hour)) {
		return nil
	}
	fmt.Printf("%s: <%d, %s> => <%d, %s>\n", relayer.aggregatorAddr.String(), latestRoundAnswer, latestRoundTime, answer, now)
	priceStrs := []string{}
	for _, price := range prices {
		priceStrs = append(priceStrs, strconv.FormatFloat(price, 'f', -1, 64))
	}
	// insert into recorder
	return relayer.recorder.PutRecords(
		relayer.aggregatorAddr.String(),
		strconv.FormatUint(number+1, 10),
		[]*Record{
			{
				Aggregator: relayer.aggregatorAddr.String(),
				Answer:     uint64(answer),
				Prices:     strings.Join(priceStrs, ","),
				Number:     number + 1,
			},
		})
}

func (relayer *exchangeRelayer) Consume(ctx context.Context) error {
	recordToConfirm, err := relayer.recorder.RecordToConfirm(relayer.aggregatorAddr.String())
	if err != nil {
		return err
	}
	if recordToConfirm != nil {
		nonce, err := relayer.nonceAt(ctx)
		if err != nil {
			return err
		}
		switch _, err := relayer.targetClient.TransactionReceipt(ctx, common.HexToHash(recordToConfirm.RelayTxHash)); errors.Cause(err) {
		case nil:
			return relayer.recorder.ConfirmRecord(recordToConfirm.ID)
		case ethereum.NotFound:
			if recordToConfirm.Nonce <= nonce && recordToConfirm.CreatedAt.Add(10*time.Minute).Before(time.Now()) {
				return relayer.recorder.ResetRecord(recordToConfirm.ID)
			}
			return nil
		default:
			return err
		}
	}
	recordToRelay, err := relayer.recorder.RecordToRelay(relayer.aggregatorAddr.String())
	if err != nil {
		return err
	}
	if recordToRelay == nil {
		return nil
	}

	opts, err := relayer.transactionOpts(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("Submitting <%s, %d>...\n", recordToRelay.Aggregator, recordToRelay.Number)
	tx, err := relayer.aggregator.Submit(
		opts,
		new(big.Int).SetUint64(recordToRelay.Answer),
		uint64(recordToRelay.CreatedAt.Unix()),
	)
	if err != nil {
		return err
	}
	return relayer.recorder.SetRecordRelayTxHash(recordToRelay.ID, tx.Hash(), opts.From, tx.Nonce())
}
