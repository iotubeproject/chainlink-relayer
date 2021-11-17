package exchange

import "github.com/pkg/errors"

type Exchange interface {
	Price() (float64, error)
}

var (
	ErrFailedToQuery   = errors.New("failed to query")
	ErrUnknownExchange = errors.New("unknown exchange")
)

func NewExchange(key string, symbol string) (Exchange, error) {
	switch key {
	case "binance":
		return NewBinanceExchange(symbol), nil
	case "coinbase":
		return NewCoinbaseExchange(symbol), nil
	case "coingecko":
		return NewCoingeckoExchange(symbol), nil
	case "huobi":
		return NewHuobiExchange(symbol), nil
	case "kucoin":
		return NewKucoinExchange(symbol), nil
	default:
		return nil, errors.Wrapf(ErrUnknownExchange, "key: %s", key)
	}
}
