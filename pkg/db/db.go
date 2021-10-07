package db

import (
	"github.com/jmoiron/sqlx"
)

func Init() (*sqlx.DB, error) {
	// Initiate database MySQL
	db, err := sqlx.Open("mysql", "root:@tcp(127.0.0.1:3307)/alta_project?parseTime=true")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
