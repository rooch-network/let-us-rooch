package btcapi

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"time"
)

type Price struct {
	Time int64   `json:"time"`
	USD  float64 `json:"USD"`
}

type BtcPriceResponse struct {
	Prices        []Price            `json:"prices"`
	ExchangeRates map[string]float64 `json:"exchangeRates"`
}

func (c *ApiClient) BtcUSDPrice() (float64, error) {
	res, err := c.mempoolBaseRequest(http.MethodGet, fmt.Sprintf("https://mempool.space/api/v1/historical-price?currency=USD&timestamp=%s", strconv.FormatInt(time.Now().Unix(), 10)), nil)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	data := &BtcPriceResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return 0, errors.Wrap(err, string(res))
	}
	return data.Prices[0].USD, nil
}
