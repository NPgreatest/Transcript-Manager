package route

import (
	"awesomeProject/route/score"
	"github.com/gin-gonic/gin"
	"log"
)

func Register(engine *gin.Engine) {
	RegisterScores(engine)
	RegisterUsers(engine)
}

func RegisterScores(engine *gin.Engine) {
	scoreGroup := engine.Group("/score")
	scoreGroup.POST("upload", Decorate(score.UploadScore))
	log.Println("/upload")
	scoreGroup.GET("get", Decorate(score.GetStudentScores))
	log.Println("/get")
}

func RegisterUsers(engine *gin.Engine) {
	//http.HandleFunc("/login", errWrapper(user.LoginUser))
	//log.Println("/login")

}
