// Copyright (c) 2019 IoTeX
// This is an alpha (internal) rrecorderease and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package relayer

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	New       string = ""
	Relayed          = "relayed"
	Confirmed        = "confirmed"
	Failed           = "failed"
)

type (
	Meta struct {
		SyncedHeight uint64 `gorm:"primary_key;unsigned"`
	}
	// Recorder is a logger based on sql to record exchange events
	Recorder struct {
		db *gorm.DB
	}
)

func (Meta) TableName() string {
	return "metas"
}

// NewRecorder returns a recorder for exchange
func NewRecorder(db *gorm.DB) (*Recorder, error) {
	if err := db.AutoMigrate(&Round{}, &Meta{}); err != nil {
		return nil, err
	}
	return &Recorder{db: db}, nil
}

func (recorder *Recorder) SyncedHeight() (uint64, error) {
	var height uint64
	switch result := recorder.db.Table("metas").Select("coalesce(max(synced_height), 0)").Find(&height); errors.Cause(result.Error) {
	case nil:
		return height, nil
	case gorm.ErrRecordNotFound:
		return 0, nil
	default:
		return 0, result.Error
	}
}

// AddTransfer creates a new transfer record
func (recorder *Recorder) Put(syncedHeight uint64, data []*Round) error {
	return recorder.db.Transaction(func(tx *gorm.DB) error {
		if len(data) != 0 {
			if result := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(data); result.Error != nil {
				return result.Error
			}
		}

		return tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "synced_height"}},
			DoUpdates: clause.AssignmentColumns([]string{"synced_height"}),
		}).Create(&Meta{SyncedHeight: syncedHeight}).Error
	})
}

// SetRelayTxHash updates a round with new tx hash
func (recorder *Recorder) SetRelayTxHash(id uint, txHash common.Hash) error {
	return recorder.db.Transaction(func(tx *gorm.DB) error {
		return tx.Table("rounds").Where("id = ?", id).Updates(map[string]interface{}{"status": Relayed, "relay_tx_hash": txHash.String()}).Error
	})
}

// Confirm marks a round as confirmed
func (recorder *Recorder) Confirm(id uint) error {
	return recorder.updateStatus(id, Confirmed)
}

// Fail marks a round as failed
func (recorder *Recorder) Fail(id uint) error {
	return recorder.updateStatus(id, Failed)
}

func (recorder *Recorder) RoundsToConfirm(aggregators ...string) ([]*Round, error) {
	return recorder.roundsByStatus(Relayed, aggregators...)
}

func (recorder *Recorder) RoundsToRelay(aggregators ...string) ([]*Round, error) {
	return recorder.roundsByStatus(New, aggregators...)
}

func (recorder *Recorder) roundsByStatus(status string, aggregators ...string) ([]*Round, error) {
	rounds := []*Round{}
	tx := recorder.db.Table("rounds")
	if len(aggregators) == 0 {
		tx = tx.Where("status = ?", status)
	} else {
		tx = tx.Where("status = ? AND aggregator in ?", status, aggregators)
	}
	if result := tx.Order("rounds.number asc").Find(&rounds); result.Error != nil {
		return nil, result.Error
	}

	return rounds, nil
}

func (recorder *Recorder) updateStatus(id uint, status string) error {
	return recorder.db.Transaction(func(tx *gorm.DB) error {
		return tx.Table("rounds").Where("id = ?", id).Update("status", status).Error
	})
}
