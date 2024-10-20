package btcapi

import (
	"crypto/tls"
	"fmt"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

type UnspentOutput struct {
	Outpoint *wire.OutPoint
	Output   *wire.TxOut
}

type BtcApiClient interface {
	BroadcastTx(tx *wire.MsgTx) (*chainhash.Hash, error)
	ListUnspent(address btcutil.Address) ([]*UnspentOutput, error)
}

func Request(method, baseURL, subPath string, requestBody io.Reader, bearerToken string) (b []byte, e error) {
	url := fmt.Sprintf("%s%s", baseURL, subPath)
	req, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Request")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	if bearerToken != "" {
		req.Header.Add("Authorization", "Bearer "+bearerToken)
	}

	// 创建一个自定义的 Transport，设置 InsecureSkipVerify
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// 使用自定义的 Transport 创建一个 http.Client
	client := &http.Client{
		Transport: transport,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send Request")
	}

	defer func(Body io.ReadCloser) {
		errC := Body.Close()
		if errC != nil {
			e = errors.Wrap(err, "failed to close response body")
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}
	return body, nil
}
