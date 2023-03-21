package DB

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DataBase struct {
	DB *sql.DB
}

func NewDataBase(confic string) *DataBase {
	confic = "user=postgres password=mypass dbname=productdb sslmode=disable"
	db, err := sql.Open("postgres", confic)
	if err != nil {

	}
	return &DataBase{DB: db}
}
