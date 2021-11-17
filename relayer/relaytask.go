// Copyright (c) 2019 IoTeX
// This is an alpha (internal) rrecorderease and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package relayer

type RelayTask struct {
	Relayer     string `gorm:"varchar(42);index;"`
	Nonce       uint64 `gorm:"index;"`
	RelayTxHash string `gorm:"varchar(66);index;"`
	Status      string `gorm:"varchar(10);not null;index;default:''"`
}
