package btcapi

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/wire"
	"io"
	"log"
)

type ApiClient struct {
	baseURL     string
	unisatURL   string
	bearerToken string
}

func NewClient(netParams *chaincfg.Params, bearerToken string) *ApiClient {
	baseURL := ""
	unisatURL := ""
	if netParams.Net == wire.MainNet {
		baseURL = "https://mempool.space/api"
		unisatURL = "https://open-api.unisat.io/v1/indexer"
	} else if netParams.Net == wire.TestNet3 {
		baseURL = "https://mempool.space/testnet/api"
		unisatURL = "https://open-api-testnet.unisat.io/v1/indexer"
	} else {
		log.Fatal("don't support other netParams")
	}
	return &ApiClient{
		baseURL:     baseURL,
		unisatURL:   unisatURL,
		bearerToken: bearerToken,
	}
}

func (c *ApiClient) mempoolRequest(method, subPath string, requestBody io.Reader) ([]byte, error) {
	return Request(method, c.baseURL, subPath, requestBody, "")
}

func (c *ApiClient) mempoolBaseRequest(method, basePath string, requestBody io.Reader) ([]byte, error) {
	return Request(method, basePath, "", requestBody, "")
}

func (c *ApiClient) unisatRequest(method, subPath string, requestBody io.Reader) ([]byte, error) {
	return Request(method, c.unisatURL, subPath, requestBody, c.bearerToken)
}
