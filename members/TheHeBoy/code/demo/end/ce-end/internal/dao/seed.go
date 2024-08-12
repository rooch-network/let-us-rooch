package dao

import (
	"gohub/internal/model"
	"gorm.io/gorm"
)

type SeedDao struct {
	BaseDao[model.SeedDO]
}

var Seed = new(SeedDao)

func (dao *SeedDao) Tx(db *gorm.DB) *SeedDao {
	return &SeedDao{BaseDao[model.SeedDO]{DB: db}}
}
