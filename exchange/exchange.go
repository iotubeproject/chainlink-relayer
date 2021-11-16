package exchange

import "errors"

type Exchange interface {
	Price() (float64, error)
}

var ErrFailedToQuery = errors.New("failed to query")
