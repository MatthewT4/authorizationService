package BLogic

import "authorizationService/internal/DB"

type BLogic struct {
	database *DB.DataBase
}

func NewBLogic(config string) *BLogic {
	return &BLogic{database: DB.NewDataBase(config)}
}
