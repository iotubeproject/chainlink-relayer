package exchange

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

type coinbaseExchange struct {
}

func NewCoinbaseExchange() Exchange {
	return &coinbaseExchange{}
}

func (*coinbaseExchange) Price() (float64, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.coinbase.com/v2/prices/IOTX-USD/spot", nil)
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

	values := map[string]map[string]string{}
	if err := json.Unmarshal(respBody, &values); err != nil {
		return 0, errors.Wrapf(err, "invalid response format %+v", err)
	}

	return strconv.ParseFloat(values["data"]["amount"], 64)
}
