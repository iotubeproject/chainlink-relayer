// Copyright (c) 2019 IoTeX
// This is an alpha (internal) rrecorderease and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package relayer

import (
	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	RelayTask
	// Target Aggregator Address
	Aggregator string `gorm:"uniqueIndex:idx_records_aggregator_round;index;not null;size:42;"`
	Number     uint64 `gorm:"uniqueIndex:idx_records_aggregator_round;index;not null;"`
	Answer     uint64 `gorm:"not null"`
	Prices     string `gorm:"not null"`
}

func NewRecord(aggregator string, answer uint64) (*Record, error) {
	return &Record{
		Aggregator: aggregator,
		Answer:     answer,
	}, nil
}

func (Record) TableName() string {
	return "records"
}
