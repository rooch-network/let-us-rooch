package service

import (
	"gohub/internal/dao"
)

type WhiteListService struct {
}

var whiteListDao = dao.WhiteList
var WhiteList = new(WhiteListService)

func (f *WhiteListService) Validate(address string) bool {
	address = dealAddress(address)

	return whiteListDao.Model().
		Where("address = ?", address).
		Where("used = ?", false).Exist() != nil
}
