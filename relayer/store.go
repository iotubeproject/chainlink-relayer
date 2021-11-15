// Copyright (c) 2019 IoTeX
// This is an alpha (internal) rproducerease and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package relayer

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	StoreBase struct{}
	Event     struct {
		gorm.Model
		Address     string      `gorm:"varchar(42);index;not null"`
		Topic0      string      `gorm:"varchar(66);index;not null"`
		Topic1      string      `gorm:"varchar(66);index"`
		Topic2      string      `gorm:"varchar(66);index"`
		Topic3      string      `gorm:"varchar(66);index"`
		Data        string      `gorm:"not null"`
		TxHash      string      `gorm:"varchar(66)"`
		Transaction Transaction `gorm:"foreignKey:hash;references:tx_hash;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
		IndexInTx   uint        `gorm:"unsigned;index;not null"`
	}

	Transaction struct {
		Hash        string `gorm:"primarykey;varchar(66);index;not null"`
		BlockNumber uint64 `gorm:"unsigned;index;not null"`
		Data        string `gorm:"not null"`
	}

	SyncedHeight struct {
		ID     string `gorm:"primarykey;varchar(20);index;not null"`
		Height uint64 `gorm:"unsigned;index;not null"`
	}
)

func (Event) TableName() string {
	return "events"
}

func (SyncedHeight) TableName() string {
	return "synced_heights"
}

func (storeBase *StoreBase) Put(
	id string,
	tx *gorm.DB,
	to uint64,
	logs []types.Log,
	txs []*types.Transaction,
) error {
	for _, t := range txs {
		if result := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(t); result.Error != nil {
			return result.Error
		}
	}
	for _, l := range logs {
		event := Event{
			Address:   l.Address.String(),
			Topic0:    l.Topics[0].String(),
			Data:      hex.EncodeToString(l.Data),
			TxHash:    l.TxHash.String(),
			IndexInTx: l.Index,
		}
		if len(l.Topics) > 1 {
			event.Topic1 = l.Topics[1].String()
			if len(l.Topics) > 2 {
				event.Topic2 = l.Topics[2].String()
				if len(l.Topics) > 3 {
					event.Topic3 = l.Topics[3].String()
				}
			}
		}
		if result := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(event); result.Error != nil {
			return result.Error
		}
	}

	return tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"height"}),
	}).Create(&SyncedHeight{ID: id, Height: to}).Error
}
