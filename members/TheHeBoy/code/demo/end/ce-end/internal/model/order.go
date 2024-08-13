package model

type OrderDO struct {
	BaseModel

	PayAddress     string  `gorm:"column:payAddress" json:"payAddress"`
	PayPrivateKey  string  `gorm:"column:payPrivateKey" json:"-"`
	Address        string  `gorm:"column:address" json:"address"`
	EstimateFee    int64   `gorm:"column:estimateFee" json:"estimateFee"`
	Fees           int64   `gorm:"column:fees" json:"fees"`
	OrderId        int64   `gorm:"column:orderId;uniqueIndex;" json:"orderId"`
	FeeRate        int64   `gorm:"column:feeRate" json:"feeRate"`
	Status         string  `gorm:"column:status" json:"status"`
	HSeed          string  `gorm:"column:hSeed;size:16;uniqueIndex" json:"hSeed"`
	RevealTxHash   string  `gorm:"column:revealTxHash" json:"revealTxHash"`
	InscriptionsId string  `gorm:"column:inscriptionsId" json:"inscriptionsId"`
	UsdPrice       float64 `gorm:"column:usdPrice" json:"usdPrice"`
	BtcPrice       int64   `gorm:"column:btcPrice" json:"btcPrice"`
	CommonTimestampsField
}

func (*OrderDO) TableName() string {
	return "orders"
}
