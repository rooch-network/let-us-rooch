package app

import (
	"github.com/thedevsaddam/govalidator"
	"gohub/internal/request/validators"
)

type OrderCreateReq struct {
	Address string `json:"address" valid:"address" form:"address"`
	FeeRate int64  `json:"feeRate" valid:"feeRate" form:"feeRate"`
}

func (r *OrderCreateReq) Validator() map[string][]string {
	rules := govalidator.MapData{
		"address": []string{"required"},
		"feeRate": []string{"required"},
	}

	messages := govalidator.MapData{
		"address": []string{
			"required:address 为必填项",
		},
		"feeRate": []string{
			"required:feeRate 为必填项",
		},
	}
	return validators.ValidateData(r, rules, messages)
}

type OrderExecuteReq struct {
	OrderId string `json:"orderId" valid:"orderId" form:"orderId"`
}

func (r *OrderExecuteReq) Validator() map[string][]string {
	rules := govalidator.MapData{
		"orderId": []string{"required"},
	}

	messages := govalidator.MapData{
		"orderId": []string{
			"required:orderId 为必填项",
		},
	}
	return validators.ValidateData(r, rules, messages)
}

type OrderListReq struct {
	Address string `json:"address" valid:"address" form:"address"`
}

func (r *OrderListReq) Validator() map[string][]string {
	rules := govalidator.MapData{
		"address": []string{"required"},
	}

	messages := govalidator.MapData{
		"address": []string{
			"required:address 为必填项",
		},
	}
	return validators.ValidateData(r, rules, messages)
}
