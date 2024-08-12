package dao

import (
	"gohub/internal/model"
	"gorm.io/gorm"
)

type OrderDao struct {
	BaseDao[model.OrderDO]
}

var Order = new(OrderDao)

func (dao *OrderDao) Tx(db *gorm.DB) *OrderDao {
	return &OrderDao{BaseDao[model.OrderDO]{DB: db}}
}
