package app

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gohub/internal/errorI"
	"gohub/internal/request/app"
	"gohub/internal/request/validators"
	"gohub/internal/service"
	"gohub/pkg/logger"
	"gohub/pkg/response"
	"strconv"
)

type OrderController struct {
}

var orderService = service.Order

func (oc *OrderController) Save(c *gin.Context) {
	req := app.OrderCreateReq{}
	if ok := validators.Validate(c, &req); !ok {
		return
	}

	orderDO, err := orderService.Save(req)

	if err != nil {
		logger.Errorv(err)
		if errors.Is(err, errorI.OrderSeedNoFind) {
			response.Error10001(c, err)
		} else {
			response.ErrorStr(c, "订单保存失败")
		}
	} else {
		response.SuccessData(c, gin.H{
			"payAddress":  orderDO.PayAddress,
			"estimateFee": orderDO.EstimateFee,
			"hSeed":       orderDO.HSeed,
			"orderId":     strconv.FormatInt(orderDO.OrderId, 10),
		})
	}
}

func (oc *OrderController) Execute(c *gin.Context) {
	req := app.OrderExecuteReq{}
	if ok := validators.Validate(c, &req); !ok {
		return
	}

	orderId, err := strconv.ParseInt(req.OrderId, 10, 64)
	if err != nil {
		response.ErrorStr(c, "订单号格式错误")
	}

	orderDO, err := orderService.ExecuteOrder(orderId)

	if err != nil {
		logger.Errorv(err)
		if errors.Is(err, errorI.OrderBalanceInsufficientError) {
			response.Error10001(c, err)
		} else if errors.Is(err, errorI.OrderNoExist) {
			response.Error10002(c, err)
		} else {
			response.ErrorStr(c, "订单执行失败")
		}
	} else {
		response.SuccessData(c, gin.H{
			"revealTxHash":   orderDO.RevealTxHash,
			"inscriptionsId": orderDO.InscriptionsId,
		})
	}
}

func (oc *OrderController) List(c *gin.Context) {
	req := app.OrderListReq{}
	if ok := validators.Validate(c, &req); !ok {
		return
	}

	list, err := orderService.List(req.Address)

	if err != nil {
		logger.Errorv(err)
		response.ErrorStr(c, "List 失败")
	} else {
		response.SuccessData(c, list)
	}
}
