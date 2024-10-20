package page

import (
	"fmt"
	"gohub/pkg/config"
)

type Req struct {
	PageNo   int      `json:"pageNo" valid:"pageNo" form:"pageNo"`
	PageSize int      `json:"pageSize" valid:"pageSize" form:"pageSize"`
	Orders   []string `json:"orders" valid:"orders" form:"orders"`
	Fields   []string `json:"fields" valid:"fields" form:"fields"`
}

func (r *Req) Validator() map[string][]string {
	return ValidatePage(r, make(map[string][]string))
}

// ValidatePage 自定义规则，验证分页参数
func ValidatePage(pageReq *Req, errs map[string][]string) map[string][]string {

	maxPageSize := config.GetInt("page.max_page_size")
	if pageReq.PageSize > maxPageSize {
		errs["pageNo"] = append(errs["pageNo"], fmt.Sprintf("页码超出最大限制%d", maxPageSize))
	}

	fields := pageReq.Fields
	orders := pageReq.Orders

	if len(fields) != len(orders) {
		errs["fields"] = append(errs["fields"], "fields 和 orders 长度不一致")
	}

	for _, field := range fields {
		if field == "" {
			errs["fields"] = append(errs["fields"], "fields 不能为空")
			break
		}
	}

	for _, order := range orders {
		if order != "asc" && order != "desc" {
			errs["orders"] = append(errs["orders"], "order 只能是 asc 或 desc")
			break
		}
	}
	return errs
}

type Resp[T any] struct {
	Total int64 `json:"total"`
	List  []T   `json:"list"` // nil 表示查询出错， [] 表示数据为空
}
