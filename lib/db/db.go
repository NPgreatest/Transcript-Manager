package db

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB(path string) error {
	db, err := sql.Open("mysql", path)
	if err != nil {
		return err
	}
	db.SetConnMaxLifetime(100)
	db.SetMaxIdleConns(10)
	if err := db.Ping(); err != nil {
		return errors.New("open database fail")
	}
	DB = db
	return nil
}
