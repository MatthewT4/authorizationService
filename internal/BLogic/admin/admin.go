package AdminBLogic

import (
	AdminDataBase "authorizationService/internal/DB/admin"
	"authorizationService/internal/Structs"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

var (
	validName    = regexp.MustCompile(`[a-zA-Zа-яА-ЯёЁ]{2,20}`)
	validSurname = regexp.MustCompile(`[a-zA-Zа-яА-ЯёЁ]{2,30}`)
	validEmail   = regexp.MustCompile(`\w{4,15}@\w{4,8}\.\w{2,5}`)
	validPhone   = regexp.MustCompile(`[0-9]{8,15}`)
	validPass    = regexp.MustCompile(`[\w*!@#$%^&?]{8,30}`)
)

type AdmBLogic struct {
	database AdminDataBase.IAdminDB
}

type IAdmBLogic interface {
	SignUp(inputUser *Structs.UserSignUpInput) (string, error)
	SignIn(inputUser *Structs.UserSignInInput) (string, error)
}

func NewAdmBLogic(config string) *AdmBLogic {
	return &AdmBLogic{database: AdminDataBase.NewAdminDB(config)}
}

func (b *AdmBLogic) SignUp(inputUser *Structs.UserSignUpInput) (string, error) {
	if !validEmail.MatchString(inputUser.Email) {
		return "invalid email", errors.New("invalid email")
	}
	if !validPass.MatchString(inputUser.Password) {
		return "invalid password", errors.New("invalid password")
	}
	if !validName.MatchString(inputUser.Name) {
		return "invalid name", errors.New("invalid name")
	}
	if !validSurname.MatchString(inputUser.Surname) {
		return "invalid surname", errors.New("invalid surname")
	}
	if !validPhone.MatchString(inputUser.Phone) {
		return "invalid phone", errors.New("invalid phone")
	}

	if err := b.database.CheckUniqUser(inputUser.Email); err != nil {
		return fmt.Sprintf("user with email %s already exists", inputUser.Email), err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(inputUser.Password), 10)
	if err != nil {
		return "error on hashing password", err
	}

	user := &Structs.User{
		Email:        inputUser.Email,
		HashPassword: string(hash),
		Name:         inputUser.Name,
		Surname:      inputUser.Surname,
		Phone:        inputUser.Phone,
	}

	if err := b.database.AddUser(user); err != nil {
		return "error on adding user to database", err
	}

	return "", nil
}

func (b *AdmBLogic) SignIn(inputUser *Structs.UserSignInInput) (string, error) {
	if !validEmail.MatchString(inputUser.Email) {
		return "invalid email", errors.New("invalid email")
	}
	if !validPass.MatchString(inputUser.Password) {
		return "invalid password", errors.New("invalid password")
	}

	return b.database.CheckPassword(inputUser)
}

//
//func CheckValid(login, password string) (int, string) {
//	if len(login) > 320 || validEmail.MatchString(login) == false {
//		return http.StatusBadRequest, ""
//	}
//
//	if len(password) > 120 || len(password) < 5 {
//		return http.StatusBadRequest, ""
//	}
//
//	jwtToken, err := AdminDataBase.GenerateJWTToken(login, password)
//	if err != nil {
//		return http.StatusInternalServerError, ""
//	}
//
//	return http.StatusOK, jwtToken
//}
