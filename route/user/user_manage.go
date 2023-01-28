package user

import (
	"awesomeProject/lib"
	"awesomeProject/response"
	"awesomeProject/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func LoginUser(ctx *gin.Context) *response.Response {
	fmt.Print("begin login...\n")
	login_info := &lib.ReqGetUser{
		Id:       ctx.Query("id"),
		Password: ctx.Query("password"),
	}
	success := service.ConfirmLogin(login_info)
	if success == true {
		fmt.Println("login sucess")
		return nil
	}
	fmt.Println("login failed")
	return nil
}
