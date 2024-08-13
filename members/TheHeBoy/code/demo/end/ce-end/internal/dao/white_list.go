package dao

import (
	"gohub/internal/model"
	"gorm.io/gorm"
)

type WhiteListDao struct {
	BaseDao[model.WhiteListDO]
}

var WhiteList = new(WhiteListDao)

func (dao *WhiteListDao) Tx(db *gorm.DB) *WhiteListDao {
	return &WhiteListDao{BaseDao[model.WhiteListDO]{DB: db}}
}
