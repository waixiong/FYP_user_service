package model

import (

	//pb "./proto"

	// "github.com/syndtr/goleveldb/leveldb"

	"getitqec.com/server/user/pkg/commons"
	"getitqec.com/server/user/pkg/dao"
)

type UserModel struct {
	UserDAO        dao.IUserDAO
	PortfolioDAO   dao.IPortfolioDAO
	StockConfigDAO dao.IStockConfigDAO
}

// InitModel ...
func InitModel(m commons.MongoDB) UserModelI {
	// dao := &dao.UserDAO{}
	_udao := dao.InitUserDAO(m)
	_pdao := dao.InitPortfolioDAO(m)
	_sdao := dao.InitStockConfigDAO(m)
	return &UserModel{UserDAO: _udao, PortfolioDAO: _pdao, StockConfigDAO: _sdao}
}
