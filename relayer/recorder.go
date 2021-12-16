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
		Key   string `gorm:"primary_key"`
		Value string `gorm:"default:0;unsigned"`
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
	if err := db.AutoMigrate(&Round{}, &Record{}, &ConfigSet{}, &Meta{}); err != nil {
		return nil, err
	}
	return &Recorder{db: db}, nil
}

func (recorder *Recorder) Value(key string) (string, error) {
	var value string
	switch result := recorder.db.Table("metas").Select("coalesce(value, '')").Where("key = ?", key).Find(&value); errors.Cause(result.Error) {
	case nil:
		return value, nil
	case gorm.ErrRecordNotFound:
		return "", nil
	default:
		return "", result.Error
	}
}

func (recorder *Recorder) PutRounds(key, value string, rounds []*Round, configs []*ConfigSet) error {
	return recorder.db.Transaction(func(tx *gorm.DB) error {
		if len(rounds) != 0 {
			if result := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(rounds); result.Error != nil {
				return result.Error
			}
		}
		if len(configs) != 0 {
			if result := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(configs); result.Error != nil {
				return result.Error
			}
		}

		return tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "key"}},
			DoUpdates: clause.AssignmentColumns([]string{"value"}),
		}).Create(&Meta{Key: key, Value: value}).Error
	})
}

func (recorder *Recorder) PutRecords(key, value string, records []*Record) error {
	return recorder.db.Transaction(func(tx *gorm.DB) error {
		if len(records) != 0 {
			if result := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(records); result.Error != nil {
				return result.Error
			}
		}

		return tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "key"}},
			DoUpdates: clause.AssignmentColumns([]string{"value"}),
		}).Create(&Meta{Key: key, Value: value}).Error
	})
}

func (recorder *Recorder) setRelayTxHash(tableName string, id uint, txHash common.Hash, relayer common.Address, nonce uint64) error {
	return recorder.db.Transaction(func(tx *gorm.DB) error {
		return tx.Table(tableName).Where("id = ?", id).Updates(map[string]interface{}{
			"status":        Relayed,
			"relay_tx_hash": txHash.String(),
			"relayer":       relayer.String(),
			"nonce":         nonce,
		}).Error
	})
}

func (recorder *Recorder) SetConfigRelayTxHash(id uint, txHash common.Hash, relayer common.Address, nonce uint64) error {
	return recorder.setRelayTxHash("configs", id, txHash, relayer, nonce)
}

func (recorder *Recorder) SetRoundRelayTxHash(id uint, txHash common.Hash, relayer common.Address, nonce uint64) error {
	return recorder.setRelayTxHash("rounds", id, txHash, relayer, nonce)
}

func (recorder *Recorder) SetRecordRelayTxHash(id uint, txHash common.Hash, relayer common.Address, nonce uint64) error {
	return recorder.setRelayTxHash("records", id, txHash, relayer, nonce)
}

func (recorder *Recorder) ConfirmRound(id uint) error {
	return recorder.updateStatus("rounds", id, Confirmed)
}

func (recorder *Recorder) ConfirmRecord(id uint) error {
	return recorder.updateStatus("records", id, Confirmed)
}

func (recorder *Recorder) FailRound(id uint) error {
	return recorder.updateStatus("rounds", id, Failed)
}

func (recorder *Recorder) FailRecord(id uint) error {
	return recorder.updateStatus("records", id, Failed)
}

func (recorder *Recorder) ResetRound(id uint) error {
	return recorder.updateStatus("rounds", id, New)
}

func (recorder *Recorder) ResetRecord(id uint) error {
	return recorder.updateStatus("records", id, New)
}

func (recorder *Recorder) RoundToConfirm(aggregator string) (*Round, error) {
	return recorder.roundByStatus(Relayed, aggregator)
}

func (recorder *Recorder) RoundsToRelay(aggregator string) (*Round, error) {
	return recorder.roundByStatus(New, aggregator)
}

func (recorder *Recorder) ConfigToRelay(aggregator string) (*ConfigSet, error) {
	return recorder.configByStatus(New, aggregator)
}

func (recorder *Recorder) RecordToConfirm(aggregator string) (*Record, error) {
	return recorder.recordByStatus(Relayed, aggregator)
}

func (recorder *Recorder) RecordToRelay(aggregator string) (*Record, error) {
	return recorder.recordByStatus(New, aggregator)
}

func (recorder *Recorder) roundByStatus(status string, aggregators ...string) (*Round, error) {
	rounds := []*Round{}
	tx := recorder.db.Table("rounds")
	if len(aggregators) == 0 {
		tx = tx.Where("status = ?", status)
	} else {
		tx = tx.Where("status = ? AND aggregator in ?", status, aggregators)
	}
	switch result := tx.Limit(1).Order("created_at ASC, number ASC").Find(&rounds); errors.Cause(result.Error) {
	case nil:
		if len(rounds) > 0 {
			return rounds[0], nil
		}
		return nil, nil
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		return nil, result.Error
	}
}

func (recorder *Recorder) recordByStatus(status string, aggregators ...string) (*Record, error) {
	records := []*Record{}
	tx := recorder.db.Table("records")
	if len(aggregators) == 0 {
		tx = tx.Where("status = ?", status)
	} else {
		tx = tx.Where("status = ? AND aggregator in ?", status, aggregators)
	}
	switch result := tx.Limit(1).Order("created_at ASC, number ASC").Find(&records); errors.Cause(result.Error) {
	case nil:
		if len(records) > 0 {
			return records[0], nil
		}
		return nil, nil
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		return nil, result.Error
	}
}

func (recorder *Recorder) configByStatus(status string, aggregators ...string) (*ConfigSet, error) {
	configs := []*ConfigSet{}
	tx := recorder.db.Table("configs")
	if len(aggregators) == 0 {
		tx = tx.Where("status = ?", status)
	} else {
		tx = tx.Where("status = ? AND aggregator in ?", status, aggregators)
	}
	switch result := tx.Limit(1).Order("created_at ASC, config_count ASC").Find(&configs); errors.Cause(result.Error) {
	case nil:
		if len(configs) > 0 {
			return configs[0], nil
		}
		return nil, nil
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		return nil, result.Error
	}
}

func (recorder *Recorder) updateStatus(tableName string, id uint, status string) error {
	return recorder.db.Transaction(func(tx *gorm.DB) error {
		return tx.Table(tableName).Where("id = ?", id).Update("status", status).Error
	})
}
