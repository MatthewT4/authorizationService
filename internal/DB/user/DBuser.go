package DataBase

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type UserDB struct {
	db *sql.DB
}

type IUserDB interface {
}

func NewDataBase(config string) *UserDB {
	db, err := sql.Open("postgres", config)
	if err != nil {

	}
	return &UserDB{db: db}
}
