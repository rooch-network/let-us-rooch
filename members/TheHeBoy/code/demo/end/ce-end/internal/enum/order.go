package enum

type OrderStatusEnum BaseEnum

var (
	OrderStatusWaitPay  = OrderStatusEnum{"1", "待支付"}
	OrderStatusComplete = OrderStatusEnum{"2", "完成"}
)
