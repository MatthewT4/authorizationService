package AdminBLogic

import (
	AdminDataBase "authorizationService/internal/DB/admin"
	"net/http"
	"regexp"
)

var validEmail = regexp.MustCompile(`\w+@\w+\.\w`)

type AdmBLogic struct {
	database AdminDataBase.IAdminDB
}

type IAdmBLogic interface{}

func NewAdmBLogic(config string) *AdmBLogic {
	return &AdmBLogic{database: AdminDataBase.NewAdminDB(config)}
}

func CheckValid(login, password string) (int, string) {
	if len(login) > 320 || validEmail.MatchString(login) == false {
		return http.StatusBadRequest, ""
	}

	if len(password) > 120 || len(password) < 5 {
		return http.StatusBadRequest, ""
	}

	jwtToken, err := AdminDataBase.GenerateJWTToken(login, password)
	if err != nil {
		return http.StatusInternalServerError, ""
	}

	return http.StatusOK, jwtToken
}
