package score

import (
	"awesomeProject/entities"
	"awesomeProject/lib"
	"awesomeProject/response"
	"awesomeProject/service"
	"awesomeProject/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func GetStudentGPA(ctx *gin.Context) *response.Response {
	token := ctx.GetHeader("Authorization")
	uid, _ := utils.VerifyToken(token)
	res, err := service.GetGPAGroup(uid, ctx.Query("alg"))
	if err != nil {
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(res)
}

func UploadScore(ctx *gin.Context) *response.Response {
	fmt.Print("begin upload\n")
	token := ctx.GetHeader("Authorization")
	uid, _ := utils.VerifyToken(token)
	credit, cerr := strconv.ParseFloat(ctx.Query("credit"), 64)
	if cerr != nil {
		return response.ResponseQueryFailed()
	}
	score, cerr := strconv.ParseFloat(ctx.Query("score"), 64)
	login_info := &lib.ReqGetScore{
		Uid:    uid,
		Name:   ctx.Query("name"),
		Credit: credit,
		Score:  score,
	}
	insert := make([]entities.Scores, 0, 0)
	insert = append(insert, entities.Scores{login_info.Uid, utils.GenerateId(1), login_info.Name, login_info.Credit, login_info.Score, 0, 0, 0})
	service.InsertScoresSql(insert)
	return response.ResponseOperateSuccess()
}

func UploadScoreFile(ctx *gin.Context) *response.Response {
	token := ctx.GetHeader("Authorization")
	uid, _ := utils.VerifyToken(token)
	file, err := ctx.FormFile("filename")
	if err != nil {
		return response.ResponseQueryFailed()
	}
	error := service.SaveScores(file, uid)
	if error != nil {
		return response.ResponseQueryFailed()
	}
	return response.ResponseOperateSuccess()
}

func GetStudentScores(ctx *gin.Context) *response.Response {
	fmt.Printf("begin getting score\n")
	token := ctx.GetHeader("Authorization")
	uid, _ := utils.VerifyToken(token)
	scores, SelectStudentScoreErr := service.GetScoresSql(uid)
	if SelectStudentScoreErr != nil {
		log.Println(SelectStudentScoreErr.Error())
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(scores)
}

func GetAllAlgorithm(ctx *gin.Context) *response.Response {
	fmt.Printf("Getting all algorithm...")
	res := service.GetAlgorithm("")
	return response.ResponseQuerySuccess(res)
}
