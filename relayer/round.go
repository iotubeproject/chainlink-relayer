// Copyright (c) 2019 IoTeX
// This is an alpha (internal) rrecorderease and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package relayer

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/iotexproject/chainlink-relayer/contract"
	"gorm.io/gorm"
)

type Round struct {
	gorm.Model
	RelayTask
	// Source Aggregator
	Aggregator string `gorm:"uniqueIndex:idx_aggregator_round;index;size:42;not null;"`
	Number     uint64 `gorm:"uniqueIndex:idx_aggregator_round;index;unsigned;not null;"`
	Ts         time.Time
	Report     string `gorm:"not null"`
	Rs         string `gorm:"not null"`
	Ss         string `gorm:"not null"`
	Vs         string `gorm:"not null"`
	TxHash     string `gorm:"varchar(66);not null;index;"`
}

func NewRound(log types.Log, tx *types.Transaction) (*Round, error) {
	event := new(contract.AggregatorNewRound)
	fmt.Println("log.Data", log.Data)
	if err := aggregatorABI.UnpackIntoInterface(event, EventNewRound, log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range aggregatorABI.Events[EventNewRound].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	fmt.Println("log.Topics", log.Topics)
	if err := abi.ParseTopics(event, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	var report, rs, ss []byte
	var vs [32]byte
	data := tx.Data()
	if bytes.HasPrefix(data, forwarderABI.Methods[MethodForward].ID) {
		res, err := forwarderABI.Methods[MethodForward].Inputs.Unpack(data[4:])
		if err != nil {
			return nil, err
		}
		data = res[1].([]byte)
		if !bytes.HasPrefix(data, ocr2aggregatorABI.Methods[MethodTransmit].ID) {
			return nil, fmt.Errorf("invalid data")
		}
		res, err = ocr2aggregatorABI.Methods[MethodTransmit].Inputs.Unpack(data[4:])
		if err != nil {
			return nil, err
		}
		report = res[1].([]byte)
		rs, err = json.Marshal(res[2].([][32]byte))
		if err != nil {
			return nil, err
		}
		ss, err = json.Marshal(res[3].([][32]byte))
		if err != nil {
			return nil, err
		}
		vs = res[4].([32]byte)
	} else {
		if !bytes.HasPrefix(data, aggregatorABI.Methods[MethodTransmit].ID) {
			return nil, fmt.Errorf("invalid data")
		}
		res, err := aggregatorABI.Methods[MethodTransmit].Inputs.Unpack(data[4:])
		if err != nil {
			return nil, err
		}
		report = res[0].([]byte)
		rs, err = json.Marshal(res[1].([][32]byte))
		if err != nil {
			return nil, err
		}
		ss, err = json.Marshal(res[2].([][32]byte))
		if err != nil {
			return nil, err
		}
		vs = res[3].([32]byte)
	}
	return &Round{
		Aggregator: log.Address.String(),
		Number:     event.RoundId.Uint64(),
		Ts:         time.Unix(event.StartedAt.Int64(), 0),
		Report:     hex.EncodeToString(report),
		Rs:         hex.EncodeToString(rs),
		Ss:         hex.EncodeToString(ss),
		Vs:         hex.EncodeToString(vs[:]),
		TxHash:     tx.Hash().String(),
	}, nil
}

func (Round) TableName() string {
	return "rounds"
}

func (round *Round) FormatData() (report []byte, rs [][32]byte, ss [][32]byte, vs [32]byte, err error) {
	report, err = hex.DecodeString(round.Report)
	if err != nil {
		return
	}
	rawRs, err := hex.DecodeString(round.Rs)
	if err != nil {
		return
	}
	if err = json.Unmarshal(rawRs, &rs); err != nil {
		return
	}
	rawSs, err := hex.DecodeString(round.Ss)
	if err != nil {
		return
	}
	if err = json.Unmarshal(rawSs, &ss); err != nil {
		return
	}
	rawVs, err := hex.DecodeString(round.Vs)
	if err != nil {
		return
	}
	copy(vs[:], rawVs)
	return
}
