package service

import (
	"encoding/hex"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/wire"
	"github.com/pkg/errors"
	"gohub/internal/dao"
	"gohub/internal/enum"
	"gohub/internal/errorI"
	"gohub/internal/model"
	"gohub/internal/ord"
	"gohub/internal/request/app"
	"gohub/pkg/btcapi"
	"gohub/pkg/config"
	"gohub/pkg/lockP"
	"gohub/pkg/logger"
	"gohub/pkg/page"
	"gohub/pkg/rooch"
	"gohub/pkg/snowflakeP"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

type OrderService struct {
}

var orderDao = dao.Order

var Order = new(OrderService)

var keyLock = lockP.NewSafeLocks()

func (s *OrderService) Save(req app.OrderCreateReq) (*model.OrderDO, error) {
	req.Address = strings.ToLower(req.Address)

	keyLock.Lock(req.Address)
	defer keyLock.Unlock(req.Address)

	hSeed := Seed.UsedTempSeed(req.Address)
	if hSeed == "" {
		return nil, errors.WithStack(errorI.OrderSeedNoFind)
	}

	// 检查是更新 hSeed 还是创建新的 order
	if orderDO := orderDao.Model().
		Where("status = ?", enum.OrderStatusWaitPay.Code).
		Where("address = ?", req.Address).Exist(); orderDO != nil {

		fileData, err := s.fillTemplate(hSeed)
		if err != nil {
			return nil, err
		}

		// 重新生成私钥
		if orderDO.HSeed != hSeed {
			logger.Infof("update hSeed before: %+v", orderDO)
			orderDO.HSeed = hSeed

			privateKey, taprootAddress, err := s.updateHSeed(fileData)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			orderDO.PayAddress = taprootAddress.EncodeAddress()
			orderDO.PayPrivateKey = hex.EncodeToString(privateKey.Serialize())
		}

		// 重新评估费用
		if orderDO.FeeRate != req.FeeRate {
			orderDO.FeeRate = req.FeeRate
			utxoPrivateKeyBytes, err := hex.DecodeString(orderDO.PayPrivateKey)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			privateKey, _ := btcec.PrivKeyFromBytes(utxoPrivateKeyBytes)
			estimateFee, err := s.estimateFee(privateKey, req.Address, fileData, req.FeeRate)
			if err != nil {
				return nil, err
			}
			orderDO.EstimateFee = estimateFee
		}

		if err := orderDao.New().Save(orderDO).Error; err != nil {
			return nil, errors.WithStack(err)
		}
		return orderDO, nil
	}

	fileData, err := s.fillTemplate(hSeed)
	if err != nil {
		return nil, err
	}

	// 生成私钥和地址
	privateKey, taprootAddress, err := s.updateHSeed(fileData)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	estimateFee, err := s.estimateFee(privateKey, req.Address, fileData, req.FeeRate)
	if err != nil {
		return nil, err
	}

	orderDO := &model.OrderDO{
		PayAddress:    taprootAddress.EncodeAddress(),
		PayPrivateKey: hex.EncodeToString(privateKey.Serialize()),
		FeeRate:       req.FeeRate,
		EstimateFee:   estimateFee,
		HSeed:         hSeed,
		Address:       req.Address,
		OrderId:       snowflakeP.Node.Generate().Int64(),
		Status:        enum.OrderStatusWaitPay.Code,
	}

	if err := orderDao.New().Create(orderDO).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return orderDO, nil
}

func (s *OrderService) ExecuteOrder(orderId int64) (*model.OrderDO, error) {
	// 检查订单是否存在
	orderDO := orderDao.Model().Where("orderId = ?", orderId).Exist()
	if orderDO == nil {
		return nil, errors.WithStack(errorI.OrderNoExist)
	}

	if orderDO.Status != enum.OrderStatusWaitPay.Code {
		return nil, errors.New("order status error, order require status is wait pay")
	}

	keyLock.Lock(orderDO.Address)
	defer keyLock.Unlock(orderDO.Address)

	utxoPrivateKeyBytes, err := hex.DecodeString(orderDO.PayPrivateKey)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	utxoPrivateKey, _ := btcec.PrivKeyFromBytes(utxoPrivateKeyBytes)
	utxoTaprootAddress, err := btcutil.DecodeAddress(orderDO.PayAddress, btcapi.NetParams)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	time.Sleep(5 * time.Second)
	// 每十秒查询一次，最多查询4次
	totalValue := int64(0)
	count := 0
	txOutPointList := make([]*wire.OutPoint, 0)
	txOutList := make([]*wire.TxOut, 0)
	for count < 4 {
		count++
		totalValue = int64(0)
		txOutPointList = txOutPointList[:0]
		txOutList = txOutList[:0]
		unspentList, err := btcapi.Client.ListUnspent(utxoTaprootAddress)
		if err != nil {
			return nil, err
		}
		for i := range unspentList {
			txOutPointList = append(txOutPointList, unspentList[i].Outpoint)
			txOutList = append(txOutList, unspentList[i].Output)
			totalValue += unspentList[i].Output.Value
		}
		if totalValue < orderDO.EstimateFee {
			time.Sleep(10 * time.Second)
		} else {
			break
		}
	}

	if totalValue < orderDO.EstimateFee {
		return nil, errors.WithStack(errorI.OrderBalanceInsufficientError)
	}

	fileData, err := s.fillTemplate(orderDO.HSeed)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	request := ord.InscriptionRequest{
		TxOutPointList: txOutPointList,
		TxOutList:      txOutList,
		TxPrivateKey:   utxoPrivateKey,
		FeeRate:        orderDO.FeeRate,
		ChargeFee:      s.getChargeFee(orderDO.Address),
		Data: ord.InscriptionData{
			ContentType: http.DetectContentType(fileData),
			Body:        fileData,
			Destination: orderDO.Address,
		},
	}

	tool, err := ord.NewInscriptionTool(&request)
	if err != nil {
		return nil, err
	}

	revealTxHash, inscriptionId, fee, err := tool.Inscribe()
	if err != nil {
		logger.Errorv(err)
		// try again
		time.Sleep(5 * time.Second)
		revealTxHash, inscriptionId, fee, err = tool.Inscribe()
		if err != nil {
			return nil, err
		}
	}

	orderDO.RevealTxHash = revealTxHash.String()
	orderDO.InscriptionsId = inscriptionId
	orderDO.Fees = fee
	orderDO.Status = enum.OrderStatusComplete.Code
	orderDO.BtcPrice = s.getChargeFee(orderDO.Address)

	usd, err := btcapi.Client.BtcUSDPrice()
	if err != nil {
		return nil, err
	}
	orderDO.UsdPrice = usd * float64(orderDO.BtcPrice) / 1e8
	err = dao.Transaction(func(tx *gorm.DB) error {
		if err := orderDao.Tx(tx).New().Save(orderDO).Error; err != nil {
			return errors.WithStack(err)
		}

		// 更新白名单
		if whiteListDO := dao.WhiteList.Tx(tx).Model().
			Where("address = ?", orderDO.Address).
			Where("used = ?", false).Exist(); whiteListDO != nil {
			whiteListDO.OrderId = orderDO.OrderId
			whiteListDO.Used = true
			if err := dao.WhiteList.Tx(tx).New().Save(whiteListDO).Error; err != nil {
				return errors.WithStack(err)
			}
		}

		err := Seed.useSeed(tx, orderDO.HSeed, orderDO.Address)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return orderDO, nil
}

func (s *OrderService) PageOrder(req page.Req) (*page.Resp[model.OrderDO], error) {
	return orderDao.Model().Order("id asc").Where("status = ?", enum.OrderStatusComplete.Code).Page(req)
}

type OrderListResp struct {
	Num       uint64 `json:"num"`
	BitcoinTx string `json:"bitcoin_tx"`
	HSeed     string `json:"hSeed"`
	//IsOpen    bool   `json:"isOpen"`
	ObjectId string `json:"objectId"`
}

func (s *OrderService) List(address string) ([]OrderListResp, error) {
	address = strings.ToLower(address)

	rows, err := seedDao.Model().Select("hSeed", "id").Where("address = ?", address).Rows()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	rowMap, err := dao.MapRows[string, uint64](rows)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	data, err := rooch.BtcQueryInscriptions(address)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	orderList := make([]model.OrderDO, 0)
	orderDao.Model().Order("id asc").
		Where("status = ?", enum.OrderStatusComplete.Code).
		Where("address = ?", address).
		Find(&orderList)

	resp := make([]OrderListResp, 0)
	for i := range orderList {
		order := orderList[i]
		resp = append(resp, OrderListResp{
			Num:       rowMap[order.HSeed],
			BitcoinTx: order.RevealTxHash,
			ObjectId:  data[order.RevealTxHash],
			HSeed:     order.HSeed,
		})
	}

	return resp, nil
}

func (s *OrderService) updateHSeed(fileData []byte) (*btcec.PrivateKey, *btcutil.AddressTaproot, error) {
	return ord.CreateAccount(btcapi.NetParams, ord.InscriptionData{
		ContentType: http.DetectContentType(fileData),
		Body:        fileData,
		Destination: "",
	})
}

func (s *OrderService) fillTemplate(hSeed string) ([]byte, error) {
	filePath := config.GetString("template_path")
	return Seed.FillTemplate(filePath, hSeed)
}

func (s *OrderService) estimateFee(privateKey *btcec.PrivateKey, receiveAddress string, fileData []byte, feeRate int64) (int64, error) {
	// mock one unspent utxo
	txOutPointList := []*wire.OutPoint{{
		Hash:  [32]byte{},
		Index: 0,
	}}
	txOutList := []*wire.TxOut{{
		PkScript: make([]byte, 32),
		Value:    1e6 - 1,
	}}

	request := ord.InscriptionRequest{
		TxOutPointList: txOutPointList,
		TxOutList:      txOutList,
		TxPrivateKey:   privateKey,
		FeeRate:        feeRate,
		ChargeFee:      s.getChargeFee(receiveAddress),
		Data: ord.InscriptionData{
			ContentType: http.DetectContentType(fileData),
			Body:        fileData,
			Destination: receiveAddress,
		},
	}

	tool, err := ord.NewInscriptionTool(&request)
	if err != nil {
		return 0, err
	}

	return tool.EstimateFee(), nil
}

func (s *OrderService) getChargeFee(address string) int64 {
	// check whitelist
	if WhiteList.Validate(address) {
		return 0
	}

	if config.GetInt64("service_fee.amount") > 0 {
		return config.GetInt64("service_fee.amount")
	}

	return 0
}
