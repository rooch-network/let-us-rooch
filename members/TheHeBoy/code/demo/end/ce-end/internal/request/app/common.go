package app

import (
	"github.com/thedevsaddam/govalidator"
	"gohub/internal/request/validators"
)

type AddressReq struct {
	Address string `json:"address" valid:"address" form:"address"`
}

func (r *AddressReq) Validator() map[string][]string {
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
