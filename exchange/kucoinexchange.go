package exchange

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

type kucoinExchange struct {
}

func NewKucoinExchange() Exchange {
	return &kucoinExchange{}
}

func (*kucoinExchange) Price() (float64, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.kucoin.com/api/v1/market/orderbook/level1?symbol=IOTX-USDT", nil)
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
	rawMessages := map[string]json.RawMessage{}
	if err := json.Unmarshal(respBody, &rawMessages); err != nil {
		return 0, errors.Wrapf(err, "invalid response format %s", respBody)
	}
	data := map[string]string{}
	if err := json.Unmarshal(rawMessages["data"], &rawMessages); err != nil {
		return 0, errors.Wrapf(err, "invalid response format %s", rawMessages["tick"])
	}

	return strconv.ParseFloat(data["price"], 64)
}
