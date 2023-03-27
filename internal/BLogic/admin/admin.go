package AdminBLogic

import (
	AdminDataBase "authorizationService/internal/DB/admin"
)

type AdmBLogic struct {
	database AdminDataBase.IAdminDB
}

type IAdmBLogic interface {
}

func NewAdmBLogic(config string) *AdmBLogic {
	return &AdmBLogic{database: AdminDataBase.NewAdminDB(config)}
}
