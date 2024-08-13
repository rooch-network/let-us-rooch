package model

type SeedDO struct {
	BaseModel

	Address string `gorm:"column:address" json:"address"`
	HSeed   string `gorm:"column:hSeed" json:"hSeed"`

	CommonTimestampsField
}

func (*SeedDO) TableName() string {
	return "seeds"
}
