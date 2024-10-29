// Copyright (c) 2019 IoTeX
// This is an alpha (internal) rrecorderease and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package relayer

import (
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/iotexproject/chainlink-relayer/contract"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type (
	Transmission struct {
		gorm.Model
		RelayTask
		// Source Aggregator
		Aggregator            string `gorm:"uniqueIndex:idx_aggregator_transmission;index;size:42;not null;"`
		AggregatorRoundId     uint32 `gorm:"uniqueIndex:idx_aggregator_transmission;index;unsigned;not null;"`
		Answer                string
		Transmitter           string
		ObservationsTimestamp uint32
		TransmissionTimestamp uint32
		Observations          string
		Observers             string
		EpochAndRound         string
		TxHash                string `gorm:"varchar(66);not null;index;"`
	}
)

func NewTransmission(log types.Log, blocktime uint64) (*Transmission, error) {
	event := new(contract.OCR2AggregatorNewTransmission)
	if err := ocr2aggregatorABI.UnpackIntoInterface(event, EventNewTransmission, log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range ocr2aggregatorABI.Events[EventNewTransmission].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(event, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}

	observers := hex.EncodeToString(event.Observers)
	observations := []string{}
	for _, observation := range event.Observations {
		observations = append(observations, observation.String())
	}

	return &Transmission{
		Aggregator:            log.Address.Hex(),
		AggregatorRoundId:     event.AggregatorRoundId,
		Answer:                event.Answer.String(),
		Transmitter:           event.Transmitter.Hex(),
		ObservationsTimestamp: event.ObservationsTimestamp,
		TransmissionTimestamp: uint32(blocktime),
		Observations:          strings.Join(observations, ","),
		Observers:             observers,
		EpochAndRound:         event.EpochAndRound.String(),
		TxHash:                log.TxHash.Hex(),
	}, nil
}

func (Transmission) TableName() string {
	return "transmissions"
}

func (tx *Transmission) FormatData() (
	aggregator common.Address,
	answer *big.Int,
	transmitter common.Address,
	observations []*big.Int,
	observers []byte,
	epochAndRound *big.Int,
	txHash common.Hash,
	err error,
) {
	aggregator = common.HexToAddress(tx.Aggregator)
	answer, success := new(big.Int).SetString(tx.Answer, 10)
	if !success {
		err = errors.Errorf("failed to parse answer %s", tx.Answer)
		return
	}
	transmitter = common.HexToAddress(tx.Transmitter)
	observations = []*big.Int{}
	for _, observation := range strings.Split(tx.Observations, ",") {
		obs, success := new(big.Int).SetString(observation, 10)
		if !success {
			err = errors.Errorf("failed to parse observation %s", observation)
			return
		}
		observations = append(observations, obs)
	}
	observers, err = hex.DecodeString(tx.Observers)
	if err != nil {
		return
	}
	epochAndRound, success = new(big.Int).SetString(tx.EpochAndRound, 10)
	if !success {
		err = errors.Errorf("failed to parse epoch and round %s", tx.EpochAndRound)
		return
	}
	txHash = common.HexToHash(tx.TxHash)
	return
}
