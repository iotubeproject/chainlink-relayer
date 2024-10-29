package relayer

import (
	"log"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/iotexproject/chainlink-relayer/contract"
	"github.com/iotexproject/chainlink-relayer/exchange"
	"github.com/iotexproject/chainlink-relayer/runner"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	ErrGasPriceTooHigh   = errors.New("Gas price is too high")
	EventNewRound        = "NewRound"
	EventNewTransmission = "NewTransmission"
	EventConfigSet       = "ConfigSet"
	MethodTransmit       = "transmit"
	MethodForward        = "forward"
	aggregatorABI        abi.ABI
	forwarderABI         abi.ABI
	ocr2aggregatorABI    abi.ABI
)

func init() {
	var err error
	aggregatorABI, err = abi.JSON(strings.NewReader(contract.AggregatorABI))
	if err != nil {
		log.Panic(err)
	}
	forwarderABI, err = abi.JSON(strings.NewReader(contract.ForwarderABI))
	if err != nil {
		log.Panic(err)
	}
	ocr2aggregatorABI, err = abi.JSON(strings.NewReader(contract.OCR2AggregatorABI))
	if err != nil {
		log.Panic(err)
	}
}

func NewService(
	interval time.Duration,
	mode int,
	startHeight uint64,
	databaseURL string,
	sourceClientURL string,
	targetChainID uint32,
	targetClientURL string,
	privateKey string,
	aggregatorPairs map[string]AggregatorConfig,
	exchangeAggregators map[string]map[string]string,
	hookUrl string,
	batchSize uint64,
) (*runner.ProducerConsumerRunner, error) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}
	recorder, err := NewRecorder(db)
	if err != nil {
		return nil, err
	}
	log.Printf("SOURCE CLIENT URL: %s\n", sourceClientURL)
	sourceClient, err := ethclient.Dial(sourceClientURL)
	if err != nil {
		return nil, err
	}
	log.Printf("TARGET CLIENT URL: %s\n", targetClientURL)
	targetClient, err := ethclient.Dial(targetClientURL)
	if err != nil {
		return nil, err
	}
	relayers := []Relayer{}
	for aggregatorAddr, exchanges := range exchangeAggregators {
		exs := []exchange.Exchange{}
		for key, symbol := range exchanges {
			ex, err := exchange.NewExchange(key, symbol)
			if err != nil {
				return nil, err
			}
			exs = append(exs, ex)
		}
		relayer, err := NewExchangeRelayer(privateKey, exs, recorder, common.HexToAddress(aggregatorAddr), targetChainID, targetClient)
		if err != nil {
			return nil, err
		}
		relayers = append(relayers, relayer)
	}
	if batchSize == 0 {
		batchSize = 99
	}
	relayer, err := NewContractRelayer(privateKey, mode, startHeight, recorder, aggregatorPairs, sourceClient, targetChainID, targetClient, hookUrl, batchSize)
	if err != nil {
		return nil, err
	}
	return runner.NewProducerConsumerRunner(interval, NewRelayerMaster(append(relayers, relayer)), func(err error) {
		log.Printf("Something goes wrong %+v\n", err)
	})
}
