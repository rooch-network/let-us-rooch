package model

type WhiteListDO struct {
	BaseModel

	Address string `gorm:"column:address" json:"address"`
	OrderId int64  `gorm:"column:orderId" json:"orderId"`
	Used    bool   `gorm:"column:used;default:false" json:"used"`
}

func (*WhiteListDO) TableName() string {
	return "whiteLists"
}
