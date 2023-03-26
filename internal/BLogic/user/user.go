package UserBLogic

import (
	"authorizationService/internal/DB/user"
)

type BLogic struct {
	database DataBase.IUserDB
}

func NewBLogic(config string) *BLogic {
	return &BLogic{database: DataBase.NewDataBase(config)}
}
