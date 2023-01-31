package route

import (
	"awesomeProject/route/score"
	"awesomeProject/route/user"
	"github.com/gin-gonic/gin"
	"log"
)

func Register(engine *gin.Engine) {
	RegisterScores(engine)
	RegisterUsers(engine)
}

func RegisterScores(engine *gin.Engine) {
	scoreGroup := engine.Group("/score")
	scoreGroup.POST("/upload", Decorate(score.UploadScore))
	scoreGroup.POST("/uploadfromfile", Decorate(score.UploadScoreFile))
	scoreGroup.GET("/get", Decorate(score.GetStudentScores))
	scoreGroup.GET("/getalg", Decorate(score.GetAllAlgorithm))
}

func RegisterUsers(engine *gin.Engine) {
	userGroup := engine.Group("/user")
	userGroup.GET("login", Decorate(user.LoginUser))
	log.Println("/login")

}
