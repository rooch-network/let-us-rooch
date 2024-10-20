package btcapi

import (
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"gohub/pkg/config"
)

var (
	NetParams     *chaincfg.Params
	Client        *ApiClient
	ChargeAddress btcutil.Address
)

func InitBtc() {
	mode := config.Get("bit.mode")
	NetParams = &chaincfg.TestNet3Params
	if mode == "mainnet" {
		NetParams = &chaincfg.MainNetParams
	}
	Client = NewClient(NetParams, config.Get("unisat_api_key"))

	addressStr := config.Get("service_fee.receive_address")

	var err error
	ChargeAddress, err = btcutil.DecodeAddress(addressStr, NetParams)
	if err != nil {
		panic(err)
	}
}
