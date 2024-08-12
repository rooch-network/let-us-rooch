package btcapi

import (
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

func (c *ApiClient) LastBlockHeight() (uint64, error) {
	res, err := c.mempoolRequest(http.MethodGet, "/blocks/tip/height", nil)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	// 将字符串转换为 uint64
	value, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return value, nil
}
