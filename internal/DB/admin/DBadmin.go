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

var database AdminDB

type User struct {
	id       int
	email    string
	password string
	name     string
	surname  string
	phone    string
}

type IAdminDB interface{}

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
	database = AdminDB{db: db}
	return &AdminDB{db: db}
}

func GenerateJWTToken(login, password string) (string, error) {
	if database.checkPassword(login, password) == false {
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

func (a AdminDB) checkPassword(login, password string) bool {
	row := a.db.QueryRow(`select * from "users" where email=$1`, login)
	user := User{}
	err := row.Scan(&user.id, &user.email, &user.password, &user.name, &user.surname, &user.phone)
	if err != nil {
		panic(err)
	}
	if (user == User{}) { // no user with this login
		return false
	}

	if user.password != password { // wrong password
		return false
	}
	return true
}
