package AdminDataBase

import "database/sql"

type AdminDB struct {
	db *sql.DB
}

type IAdminDB interface {
}

func NewAdminDB(config string) *AdminDB {
	db, err := sql.Open("postgres", config)
	if err != nil {

	}
	return &AdminDB{db: db}
}
