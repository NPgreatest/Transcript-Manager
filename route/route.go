package route

import (
	"awesomeProject/route/score"
	"awesomeProject/route/user"
	"github.com/gin-gonic/gin"
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
	scoreGroup.GET("/getgpa", Decorate(score.GetStudentGPA))
	scoreGroup.GET("/getalg", Decorate(score.GetAllAlgorithm))
}

func RegisterUsers(engine *gin.Engine) {
	engine.POST("/user/login", Decorate(user.LoginUser))
	userGroup := engine.Group("/user")
	userGroup.Use(user.LoginAuthenticationMiddleware())

}
