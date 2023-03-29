package AdminDataBase

import (
	"database/sql"
	"github.com/golang-jwt/jwt"
)

type AdminDB struct {
	db *sql.DB
}

type IAdminDB interface{}

type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func NewAdminDB(config string) *AdminDB {
	db, err := sql.Open("postgres", config)
	if err != nil {
	}
	return &AdminDB{db: db}
}

func GenerateJWTToken(login string) (string, error) {
	claims := Claims{
		jwt.StandardClaims{
			ExpiresAt: 86400, //так и не понял, секунды тут или чет другое...
			Issuer:    "test",
		},
		login,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	str, err := token.SignedString([]byte("AllYourBase"))
	return str, err
}
