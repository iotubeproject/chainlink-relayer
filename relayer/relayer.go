package relayer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-lark/lark"
	"github.com/pkg/errors"
)

var ErrInsufficientBalance = errors.Errorf("insuffient balance for gas fee")

type (
	Relayer interface {
		Produce(context.Context) error
		Consume(context.Context) error
	}
	abstractRelayer struct {
		notificationBot    *lark.Bot
		targetClient       *ethclient.Client
		recorder           *Recorder
		privateKey         *ecdsa.PrivateKey
		targetChainID      *big.Int
		gasPriceUpperBound *big.Int
		gasLimit           uint64
	}
	relayerMaster struct {
		relayers []Relayer
	}
)

func (relayer *abstractRelayer) alert(msg string) {
	if relayer.notificationBot == nil {
		return
	}
	_, err := relayer.notificationBot.PostNotificationV2(
		lark.NewMsgBuffer(lark.MsgText).Text(msg).Build(),
	)
	if err != nil {
		fmt.Printf("failed to send alert %+v\n", err)
	}
}

func (relayer *abstractRelayer) address() common.Address {
	return crypto.PubkeyToAddress(relayer.privateKey.PublicKey)
}

func (relayer *abstractRelayer) nonceAt(ctx context.Context) (uint64, error) {
	return relayer.targetClient.NonceAt(ctx, relayer.address(), nil)
}

func (relayer *abstractRelayer) balance(ctx context.Context) (*big.Int, error) {
	return relayer.targetClient.BalanceAt(ctx, relayer.address(), nil)
}

func (relayer *abstractRelayer) transactionOpts(ctx context.Context) (*bind.TransactOpts, error) {
	relayerAddr := relayer.address()

	opts, err := bind.NewKeyedTransactorWithChainID(relayer.privateKey, relayer.targetChainID)
	if err != nil {
		return nil, err
	}
	opts.Value = big.NewInt(0)
	opts.GasLimit = relayer.gasLimit
	gasPrice, err := relayer.targetClient.SuggestGasPrice(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get suggested gas price")
	}
	if gasPrice.Cmp(relayer.gasPriceUpperBound) >= 0 {
		return nil, errors.Wrapf(ErrGasPriceTooHigh, "suggested gas price %d > limit %d", gasPrice, relayer.gasPriceUpperBound)
	}
	opts.GasPrice = gasPrice
	balance, err := relayer.targetClient.BalanceAt(ctx, relayerAddr, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get balance of operator account")
	}
	gasFee := new(big.Int).Mul(new(big.Int).SetUint64(opts.GasLimit), opts.GasPrice)
	if gasFee.Cmp(balance) > 0 {
		return nil, ErrInsufficientBalance
	}
	nonce, err := relayer.targetClient.PendingNonceAt(ctx, relayerAddr)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to fetch pending nonce for %s", relayerAddr)
	}
	opts.Nonce = new(big.Int).SetUint64(nonce)

	return opts, nil
}

func NewRelayerMaster(relayers []Relayer) *relayerMaster {
	return &relayerMaster{
		relayers: relayers,
	}
}

func (master *relayerMaster) Consume(ctx context.Context) error {
	for _, relayer := range master.relayers {
		if err := relayer.Consume(ctx); err != nil {
			fmt.Printf("failed to consume, %+v\n", err)
		}
		time.Sleep(5 * time.Second)
	}
	return nil
}

func (master *relayerMaster) Produce(ctx context.Context) error {
	for _, relayer := range master.relayers {
		if err := relayer.Produce(ctx); err != nil {
			fmt.Printf("failed to produce, %+v\n", err)
		}
	}
	return nil
}
