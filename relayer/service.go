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
	ErrGasPriceTooHigh = errors.New("Gas price is too high")
	EventNewRound      = "NewRound"
	EventConfigSet     = "ConfigSet"
	MethodTransmit     = "transmit"
	aggregatorABI      abi.ABI
)

func init() {
	var err error
	aggregatorABI, err = abi.JSON(strings.NewReader(contract.AggregatorABI))
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
	aggregatorPairs map[string]string,
	exchangeAggregators map[string]map[string]string,
	hookUrl string,
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
	pairs := map[common.Address]common.Address{}
	for origin, shadow := range aggregatorPairs {
		pairs[common.HexToAddress(origin)] = common.HexToAddress(shadow)
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
	relayer, err := NewContractRelayer(privateKey, mode, startHeight, recorder, pairs, sourceClient, targetChainID, targetClient, hookUrl)
	if err != nil {
		return nil, err
	}
	return runner.NewProducerConsumerRunner(interval, NewRelayerMaster(append(relayers, relayer)), func(err error) {
		log.Printf("Something goes wrong %+v\n", err)
	})
}
