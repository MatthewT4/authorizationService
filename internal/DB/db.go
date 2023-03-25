package DB

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DataBase struct {
	DB *sql.DB
}

func NewDataBase(config string) *DataBase {
	//config = "user=postgres password=mypass dbname=productdb sslmode=disable"
	db, err := sql.Open("postgres", config)
	if err != nil {

	}
	return &DataBase{DB: db}
}
