package errorI

import e "errors"

var OrderBalanceInsufficientError = e.New("balance insufficient")
var OrderNoExist = e.New("order no exist")
var OrderExist = e.New("oder exists, can't save again")
var OrderSeedNoFind = e.New("hSeed no find")
var OrderStopMint = e.New("stop mint")
