package exchange

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

type binanceExchange struct {
	queryUrl string
}

func NewBinanceExchange(symbol string) Exchange {
	return &binanceExchange{queryUrl: "https://api.binance.com/api/v3/ticker/price?symbol=" + symbol + "USDT"}
}

func (be *binanceExchange) Price() (float64, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", be.queryUrl, nil)
	if err != nil {
		return 0, errors.Wrap(ErrFailedToQuery, err.Error())
	}
	req.Header.Set("Accepts", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return 0, errors.Wrap(ErrFailedToQuery, err.Error())
	}
	if resp.StatusCode != 200 {
		return 0, errors.Errorf("invalid response status %s", resp.Status)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, errors.Errorf("invalid response format %+v", err)
	}

	values := map[string]string{}
	if err := json.Unmarshal(respBody, &values); err != nil {
		return 0, errors.Wrapf(err, "invalid response format %s", string(respBody))
	}

	return strconv.ParseFloat(values["price"], 64)
}
