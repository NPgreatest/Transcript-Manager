package user

import (
	"awesomeProject/lib"
	"awesomeProject/response"
	"awesomeProject/service"
	"awesomeProject/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func LoginUser(ctx *gin.Context) *response.Response {
	fmt.Print("begin login...\n")
	login_info := &lib.ReqGetUser{
		Id:       ctx.Query("id"),
		Password: ctx.Query("password"),
	}
	success := service.ConfirmLogin(login_info)
	if success == false {
		fmt.Println("login failed")
		return response.NewResponseOkND(response.LoginFailed)
	}
	fmt.Println("login success")
	token, err := utils.CreateToken(login_info.Id, time.Hour*24)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return response.NewResponseOkND(response.OperateFailed)
	}
	return response.NewResponseOk(response.LoginSuccess, token, login_info.Id)
}

func LoginAuthenticationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			fmt.Println("未获得授权, ip:%s", ctx.Request.RemoteAddr)
			ctx.JSON(http.StatusOK, &(response.NewResponseOkND(response.Unauthorized).R))
			ctx.Abort()
			return
		}

		if _, ok := utils.VerifyToken(token); !ok {
			ctx.JSON(http.StatusOK, &(response.NewResponseOkND(response.Unauthorized).R))
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
