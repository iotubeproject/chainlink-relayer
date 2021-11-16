package exchange

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type coingeckoExchange struct {
}

func NewCoingeckoExchange() Exchange {
	return &coingeckoExchange{}
}

func (*coingeckoExchange) Price() (float64, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.coingecko.com/api/v3/simple/price?ids=iotex&vs_currencies=USD", nil)
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
	values := map[string]map[string]float64{}
	if err := json.Unmarshal(respBody, &values); err != nil {
		return 0, errors.Errorf("invalid response format %+v", err)
	}

	return values["iotex"]["usd"], nil
}
