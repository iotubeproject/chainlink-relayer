package relayer

import (
	"log"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/iotexproject/chainlink-relayer/contract"
	"github.com/iotexproject/chainlink-relayer/runner"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	ErrGasPriceTooHigh = errors.New("Gas price is too high")
	EventNewRound      = "NewRound"
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
	startHeight uint64,
	databaseURL string,
	aggregatorPairs map[string]string,
	sourceClientURL string,
	targetClientURL string,
	privateKey string,
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
	sourceClient, err := ethclient.Dial(sourceClientURL)
	if err != nil {
		return nil, err
	}
	targetClient, err := ethclient.Dial(targetClientURL)
	if err != nil {
		return nil, err
	}
	pairs := map[common.Address]common.Address{}
	for origin, shadow := range aggregatorPairs {
		pairs[common.HexToAddress(origin)] = common.HexToAddress(shadow)
	}
	relayer, err := NewRelayer(privateKey, startHeight, recorder, pairs, sourceClient, targetClient)
	if err != nil {
		return nil, err
	}
	return runner.NewProducerConsumerRunner(interval, relayer, func(err error) {
		log.Printf("Something goes wrong %+v\n", err)
	})
}
