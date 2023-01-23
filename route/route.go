package route

import (
	"awesomeProject/route/score"
	"awesomeProject/route/user"
	"log"
	"net/http"
)

type appHandel func(res http.ResponseWriter, req *http.Request) error

func RegisterRoutes() {
	http.HandleFunc("/login", errWrapper(user.LoginUser))
	log.Println("/login")
	http.HandleFunc("/upload", errWrapper(score.UploadScore))
	log.Println("/upload")
	http.HandleFunc("/get", errWrapper(score.GetStudentScores))
	log.Println("/get")
}
