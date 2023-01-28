package service

import (
	"awesomeProject/entities"
	"awesomeProject/lib"
	"awesomeProject/lib/db"
	"fmt"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func ConfirmLogin(param *lib.ReqGetUser) bool {
	sql := "SELECT * FROM users where id=? AND password=?"
	prepare, Prepareerr := db.DB.Prepare(sql)
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
