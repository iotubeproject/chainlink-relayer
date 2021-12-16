// Copyright (c) 2019 IoTeX
// This is an alpha (internal) rrecorderease and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package relayer

import (
	"encoding/hex"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/iotexproject/chainlink-relayer/contract"
	"gorm.io/gorm"
)

type (
	ConfigSet struct {
		gorm.Model
		RelayTask
		// Source Aggregator
		Aggregator           string `gorm:"uniqueIndex:idx_aggregator_config_count;index;size:42;not null;"`
		ConfigCount          uint64 `gorm:"uniqueIndex:idx_aggregator_config_count;index;unsigned;not null;"`
		Signers              string
		Transmitters         string
		Threshold            uint8
		EncodedConfigVersion uint64
		Encoded              string
		TxHash               string `gorm:"varchar(66);not null;index;"`
	}
)

func NewConfigSet(log types.Log) (*ConfigSet, error) {
	event := new(contract.AggregatorConfigSet)
	if err := aggregatorABI.UnpackIntoInterface(event, EventConfigSet, log.Data); err != nil {
		return nil, err
	}
	signers := []string{}
	for _, signer := range event.Signers {
		signers = append(signers, signer.String())
	}
	transmitters := []string{}
	for _, transmitter := range event.Transmitters {
		transmitters = append(transmitters, transmitter.String())
	}

	return &ConfigSet{
		Aggregator:           log.Address.String(),
		ConfigCount:          event.ConfigCount,
		Signers:              strings.Join(signers, ","),
		Transmitters:         strings.Join(transmitters, ","),
		Threshold:            event.Threshold,
		EncodedConfigVersion: event.EncodedConfigVersion,
		Encoded:              hex.EncodeToString(event.Encoded),
		TxHash:               log.TxHash.Hex(),
	}, nil
}

func (ConfigSet) TableName() string {
	return "configs"
}

func (cs *ConfigSet) FormatData() (signers []common.Address, transmitters []common.Address, threshold uint8, configVersion uint64, encoded []byte, err error) {
	encoded, err = hex.DecodeString(cs.Encoded)
	if err != nil {
		return
	}
	for _, signer := range strings.Split(cs.Signers, ",") {
		signers = append(signers, common.HexToAddress(signer))
	}
	for _, transmitter := range strings.Split(cs.Transmitters, ",") {
		transmitters = append(transmitters, common.HexToAddress(transmitter))
	}
	threshold = cs.Threshold
	configVersion = cs.EncodedConfigVersion
	return
}
