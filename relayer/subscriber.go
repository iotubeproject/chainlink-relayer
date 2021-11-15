// Copyright (c) 2019 IoTeX
// This is an alpha (internal) rproducerease and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package relayer

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type (
	CallbackFunc func(ctx context.Context, from, to uint64, logs []types.Log, txs []*types.Transaction) error
	Subscriber   struct {
		client       *ethclient.Client
		contracts    []common.Address
		topics       [][]common.Hash
		offset       uint64
		pagination   uint16
		confirmBlock uint8
		callback     CallbackFunc
	}
)

// NewSubscriber returns a subscriber
func NewSubscriber(
	client *ethclient.Client,
	contracts []common.Address,
	topics [][]common.Hash,
	offset uint64,
	pagination uint16,
	confirmBlock uint8,
	callback CallbackFunc,
) (*Subscriber, error) {
	return &Subscriber{
		client:       client,
		offset:       offset,
		pagination:   pagination,
		confirmBlock: confirmBlock,
		contracts:    contracts,
		topics:       topics,
		callback:     callback,
	}, nil
}

var ErrNotReady = errors.New("not ready yet")

func (subscriber *Subscriber) prepareQuery(ctx context.Context) (ethereum.FilterQuery, error) {
	startHeight := subscriber.offset
	tipHeight, err := subscriber.client.BlockNumber(ctx)
	if err != nil {
		return ethereum.FilterQuery{}, err
	}
	if tipHeight <= uint64(subscriber.confirmBlock) || tipHeight-uint64(subscriber.confirmBlock) <= startHeight {
		return ethereum.FilterQuery{}, ErrNotReady
	}
	endHeight := startHeight + uint64(subscriber.pagination)
	if tipHeight < endHeight {
		endHeight = tipHeight
	}

	return ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(startHeight),
		ToBlock:   new(big.Int).SetUint64(endHeight),
		Addresses: subscriber.contracts,
		Topics:    subscriber.topics,
	}, nil
}

// Pull pulls logs and transactions
func (subscriber *Subscriber) Pull(ctx context.Context) error {
	query, err := subscriber.prepareQuery(ctx)
	if err != nil {
		return err
	}

	logs, err := subscriber.client.FilterLogs(ctx, query)
	if err != nil {
		return err
	}
	seen := map[common.Hash]interface{}{}
	txs := []*types.Transaction{}
	for _, log := range logs {
		if _, ok := seen[log.TxHash]; ok {
			continue
		}
		tx, _, err := subscriber.client.TransactionByHash(ctx, log.TxHash)
		if err != nil {
			return err
		}
		seen[log.TxHash] = struct{}{}
		txs = append(txs, tx)
	}
	if err := subscriber.callback(
		ctx,
		query.FromBlock.Uint64(),
		query.ToBlock.Uint64(),
		logs,
		txs,
	); err != nil {
		return err
	}
	subscriber.offset = query.ToBlock.Uint64() + 1
	return nil
}
