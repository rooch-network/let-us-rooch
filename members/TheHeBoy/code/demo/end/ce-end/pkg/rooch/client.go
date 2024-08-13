package rooch

import (
	"crypto/tls"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"strings"
)

const baseUrl = "https://test-seed.rooch.network"
const contractAddress = "0xdb2e764a11715cc8cd056adf867053ba022a5d06116bfcde751b8b96a7e5f978::colored_egg7"

func request(requestBody string) (g *gjson.Result, e error) {
	req, err := http.NewRequest(http.MethodPost, baseUrl, strings.NewReader(requestBody))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Request")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

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
	data := gjson.ParseBytes(body).Get("result")

	if !data.Exists() {
		return nil, errors.New(string(body))
	}

	return &data, nil
}

func BtcQueryInscriptions(address string) (map[string]string, error) {

	data := fmt.Sprintf(`{"id": 101,"jsonrpc": "2.0","method": "btc_queryInscriptions","params": [{"owner":"%s"}, null, "50", true]}`, address)

	res, err := request(data)
	if err != nil {
		return nil, err
	}

	// 初始化map
	inscriptionsMap := make(map[string]string)

	for _, result := range res.Get("data").Array() {
		bitcoinTxID := result.Get("value.bitcoin_txid").String()
		objectID := result.Get("id").String()
		inscriptionsMap[bitcoinTxID] = objectID
	}

	return inscriptionsMap, nil
}

func IsOpens() ([]bool, error) {
	data := fmt.Sprintf(`{"id": 101,"jsonrpc": "2.0","method": "rooch_executeViewFunction","params": [{"function_id":"%s::is_opens", "ty_args":[], "args":["%s"]}]}`,
		contractAddress, "0xd23c49fb9a742624498390a5b59c90e968c11d52b6437278a0e2eb9b828241037e52904587e672bb763e6ee07a313e8594e9b7a89841f7d428db87917d566b22")

	res, err := request(data)
	if err != nil {
		return nil, err
	}

	if !res.Get("return_values").Exists() {
		return nil, errors.New(res.Raw)
	}

	results := make([]bool, 0)
	for _, result := range res.Get("return_values").Array() {
		for _, r := range result.Get("decoded_value").Array() {
			results = append(results, r.Bool())
		}
	}
	return results, nil
}
