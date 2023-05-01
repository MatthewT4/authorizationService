package AdminDataBase

import (
	"database/sql"
	"errors"
	"github.com/golang-jwt/jwt"
	_ "github.com/lib/pq"
)

type AdminDB struct {
	db *sql.DB
}

type User struct {
	id       int
	email    string
	password string
	name     string
	surname  string
	phone    string
}

type IAdminDB interface {
}

type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

type Configs struct {
	ServerHost string
	ServerPort string
	PgHost     string
	PgPort     string
	PgUser     string
	PgPass     string
	PgBase     string
}

func NewAdminDB(config string) *AdminDB {
	db, err := sql.Open("postgres", config)
	if err != nil {
	}
	return &AdminDB{db: db}
}

func GenerateJWTToken(login, password string) (string, error) {
	if checkPassword(login, password) == false {
		return "", errors.New("incorrect login or password")
	}

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

func checkPassword(login, password string) bool {
	config := "host=127.0.0.1 port=5432 user=postgres password=1234 dbname=test sslmode=disable" // тут мой конфиг
	db, err := sql.Open("postgres", config)
	if err != nil {
		panic(err)
	}

	row := db.QueryRow(`select * from "user" where email=$1`, login)
	if err != nil {
		panic(err)
	}
	user := User{}
	err = row.Scan(&user.id, &user.email, &user.password, &user.name, &user.surname, &user.phone)
	if (user == User{}) { // no user with this login
		return false
	}

	if user.password != password { // wrong password
		return false
	}
	return true
}
