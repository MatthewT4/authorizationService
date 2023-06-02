package AdminDataBase

import (
	"authorizationService/internal/Structs"
	"database/sql"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

type AdminDB struct {
	db *sql.DB
}

type IAdminDB interface {
	AddUser(user *Structs.User) error
	CheckUniqUser(email string) error
	CheckPassword(user *Structs.UserSignInInput) (string, error)
}

func NewAdminDB(config string) *AdminDB {
	db, err := sql.Open("postgres", config)
	if err != nil {
		log.Fatal(err)
	}
	return &AdminDB{db: db}
}

func (d *AdminDB) CheckUniqUser(email string) error {
	var userId int

	row := d.db.QueryRow("SELECT user_id FROM \"users\" WHERE email=$1", email)
	row.Scan(&userId)
	if userId != 0 {
		return errors.New("this user already exists")
	}
	return nil
}

func (d *AdminDB) AddUser(user *Structs.User) error {
	_, err := d.db.Exec("INSERT INTO \"users\" (email,password,name,surname,phone)"+
		"VALUES ($1,$2,$3,$4,$5)", user.Email, user.HashPassword, user.Name, user.Surname, user.Phone)
	return err
}

func (d *AdminDB) CheckPassword(user *Structs.UserSignInInput) (string, error) {
	var userId int
	var hashPassword string
	row := d.db.QueryRow("SELECT user_id, password FROM \"users\" WHERE email=$1", user.Email)

	if err := row.Scan(&userId, &hashPassword); err != nil {
		return "some wrong on server", err
	}
	if hashPassword == "" {
		return "no user with this email, please signup", errors.New("no user with this email")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(user.Password)); err != nil {
		return "incorrect password", err
	}

	return d.CreateJWTToken(userId)
}

func (d *AdminDB) CreateJWTToken(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

//func GenerateJWTToken(login, password string) (string, error) {
//	if database.checkPassword(login, password) == false {
//		return "", errors.New("incorrect login or password")
//	}
//
//	claims := Claims{
//		jwt.StandardClaims{
//			ExpiresAt: 86400, //так и не понял, секунды тут или чет другое...
//			Issuer:    "test",
//		},
//		login,
//	}
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	str, err := token.SignedString([]byte("AllYourBase"))
//	return str, err
//}
//
//func (a AdminDB) checkPassword(login, password string) bool {
//	row := a.db.QueryRow(`select * from "users" where email=$1`, login)
//	user := User{}
//	err := row.Scan(&user.id, &user.email, &user.password, &user.name, &user.surname, &user.phone)
//	if err != nil {
//		panic(err)
//	}
//	if (user == User{}) { // no user with this login
//		return false
//	}
//
//	if user.password != password { // wrong password
//		return false
//	}
//	return true
//}
