package user

import (
	"awesomeProject/lib"
	"awesomeProject/service"
	"fmt"
	"net/http"
)

func LoginUser(res http.ResponseWriter, req *http.Request) error {
	fmt.Print("begin login...\n")
	login_info := &lib.ReqGetUser{
		Id:       req.FormValue("id"),
		Password: req.FormValue("password"),
	}
	userDB := service.NewUserDB()
	success := userDB.ConfirmLogin(login_info)
	if success == true {
		fmt.Println("login sucess")
		return nil
	}
	fmt.Println("login failed")
	return nil
}
