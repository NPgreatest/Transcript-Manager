package db

import (
	"errors"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sqlx.DB

func ConnectDB(path string) error {
	db, err := sqlx.Connect("mysql", path)
	if err != nil {
		return errors.New("open database fail")
	}
	db.SetConnMaxLifetime(100)
	db.SetMaxIdleConns(10)
	if err := db.Ping(); err != nil {
		return errors.New("open database fail")
	}
	DB = db
	return nil
}
