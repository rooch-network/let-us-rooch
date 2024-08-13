package hashidsP

import (
	"github.com/speps/go-hashids/v2"
)

var HashID *hashids.HashID

func InitHashIds() {
	hd := hashids.NewData()
	hd.Salt = "76faccf3211631cf8b56cd1b89e6bf28"
	hd.MinLength = 16
	var err error
	HashID, err = hashids.NewWithData(hd)
	if err != nil {
		panic(err)
	}
}
