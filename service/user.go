package service

import (
	"awesomeProject/entities"
	"awesomeProject/lib"
	"awesomeProject/lib/db"
	"database/sql"
	"fmt"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
type UserDB struct {
	User
	DB *sql.DB
}

func NewUserDB() *UserDB {
	return &UserDB{
		DB:   db.DB,
		User: User{},
	}
}

func (s *UserDB) ConfirmLogin(param *lib.ReqGetUser) bool {
	sql := "SELECT * FROM users where id=? AND password=?"
	prepare, Prepareerr := s.DB.Prepare(sql)
	if Prepareerr != nil {
		return false
	}
	row := prepare.QueryRow(param.Id, param.Password)
	res := &entities.Users{}
	row.Scan(&res.Id, &res.Name, &res.Password)
	if res.Id == "" {
		fmt.Println("mysql cannot find any data of this user")
		return false
	}
	return true
}
