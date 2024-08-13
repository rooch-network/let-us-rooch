package btcapi

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"gohub/pkg/logger"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type Utxo struct {
	Address      string `json:"address"`
	CodeType     int    `json:"codeType"`
	Height       int    `json:"height"`
	Idx          int    `json:"idx"`
	Inscriptions []any  `json:"inscriptions"`
	IsOpInRBF    bool   `json:"isOpInRBF"`
	Satoshi      int64  `json:"satoshi"`
	ScriptPk     string `json:"scriptPk"`
	ScriptType   string `json:"scriptType"`
	Txid         string `json:"txid"`
	Vout         int    `json:"vout"`
}

type Brc20Detail struct {
	Address                string `json:"address"`
	OverallBalance         string `json:"overallBalance"`
	TransferableBalance    string `json:"transferableBalance"`
	AvailableBalance       string `json:"availableBalance"`
	AvailableBalanceSafe   string `json:"availableBalanceSafe"`
	AvailableBalanceUnSafe string `json:"availableBalanceUnSafe"`
}

type Brc20PageResponse struct {
	Height int           `json:"height"`
	Total  int           `json:"total"`
	Start  int           `json:"start"`
	Detail []Brc20Detail `json:"detail"`
}

func (c *ApiClient) ListUnspent(address btcutil.Address) ([]*UnspentOutput, error) {
	res, err := c.unisatRequest(http.MethodGet, fmt.Sprintf("/address/%s/utxo-data?cursor=%d&size=%d", address.EncodeAddress(), 0, 16), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	jsonStr := string(res)
	if gjson.Get(jsonStr, "code").Int() != 0 {
		return nil, errors.New("API returned an error")
	}

	unspentOutputs := make([]*UnspentOutput, 0)
	utxos := gjson.Get(jsonStr, "data.utxo").Array()
	for _, utxo := range utxos {
		txHash, err := chainhash.NewHashFromStr(utxo.Get("txid").String())
		if err != nil {
			return nil, err
		}
		scriptPk, err := hex.DecodeString(utxo.Get("scriptPk").String())
		if err != nil {
			return nil, err
		}

		unspentOutputs = append(unspentOutputs, &UnspentOutput{
			Outpoint: wire.NewOutPoint(txHash, uint32(utxo.Get("vout").Int())),
			Output:   wire.NewTxOut(utxo.Get("satoshi").Int(), scriptPk),
		})
	}
	return unspentOutputs, nil
}

func (c *ApiClient) GetAddressByInscriptionId(inscriptionId string) (string, error) {
	res, err := c.unisatRequest(http.MethodGet, fmt.Sprintf("/inscription/info/%s", inscriptionId), nil)
	if err != nil {
		return "", errors.WithStack(err)
	}

	var resData Response
	err = json.Unmarshal(res, &resData)
	if err != nil {
		logger.Error(string(res))
		return "", errors.WithStack(err)
	}

	dataMap, ok := resData.Data.(map[string]interface{})
	if !ok {
		return "", errors.New("failed to parse data")
	}

	return dataMap["address"].(string), nil
}

func (c *ApiClient) GetBrc20Page(ticker string, start int, limit int) (*Brc20PageResponse, error) {
	res, err := c.unisatRequest(http.MethodGet, fmt.Sprintf("/brc20/%s/holders?start=%d&limit=%d", ticker, start, limit), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var resData Response
	err = json.Unmarshal(res, &resData)
	if err != nil {
		logger.Error(string(res))
		return nil, errors.WithStack(err)
	}

	var page Brc20PageResponse
	dataBytes, err := json.Marshal(resData.Data)
	if err != nil {
		logger.Error(resData.Data)
		return nil, errors.New("failed to marshal data")
	}

	err = json.Unmarshal(dataBytes, &page)
	if err != nil {
		logger.Error(string(dataBytes))
		return nil, errors.New("failed to parse data")
	}

	return &page, nil
}
