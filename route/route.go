package route

import (
	"awesomeProject/route/score"
	"awesomeProject/route/user"
	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {
	RegisterPublic(engine)
	RegisterUsers(engine)
}

func RegisterPublic(engine *gin.Engine) {
	scoreGroup := engine.Group("/all")
	scoreGroup.GET("/getalg", Decorate(score.GetAllAlgorithm))
}

func RegisterUsers(engine *gin.Engine) {
	engine.POST("/user/login", Decorate(user.LoginUser))
	userGroup := engine.Group("/user")
	userGroup.Use(user.LoginAuthenticationMiddleware())
	userGroup.POST("/upload", Decorate(score.UploadScore))
	userGroup.POST("/uploadfromfile", Decorate(score.UploadScoreFile))
	userGroup.GET("/get", Decorate(score.GetStudentScores))
	userGroup.GET("/getgpa", Decorate(score.GetStudentGPA))
}
