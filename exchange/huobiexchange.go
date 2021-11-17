package exchange

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type huobiExchange struct {
	queryUrl string
}

func NewHuobiExchange(symbol string) Exchange {
	return &huobiExchange{queryUrl: "https://api.huobi.pro/market/trade?symbol=" + symbol + "usdt"}
}

func (he *huobiExchange) Price() (float64, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", he.queryUrl, nil)
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
	if err := json.Unmarshal(rawMessages["tick"], &rawMessages); err != nil {
		return 0, errors.Wrapf(err, "invalid response format %s", rawMessages["tick"])
	}
	data := []map[string]json.RawMessage{}
	if err := json.Unmarshal(rawMessages["data"], &data); err != nil {
		return 0, errors.Wrapf(err, "invalid response format %s", rawMessages["data"])
	}
	retval := float64(0)
	for _, d := range data {
		var price float64
		if err := json.Unmarshal(d["price"], &price); err != nil {
			return 0, errors.Wrapf(err, "invalid price format %s", string(d["price"]))
		}
		retval += price
	}

	return retval / float64(len(data)), nil
}
