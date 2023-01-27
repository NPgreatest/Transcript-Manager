package score

import (
	"awesomeProject/entities"
	"awesomeProject/lib"
	"awesomeProject/response"
	"awesomeProject/service"
	"awesomeProject/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func UploadScore(ctx *gin.Context) *response.Response {
	fmt.Print("begin upload\n")
	credit, cerr := strconv.ParseFloat(ctx.Query("credit"), 64)
	if cerr != nil {
		return response.ResponseQueryFailed()
	}
	score, cerr := strconv.ParseFloat(ctx.Query("score"), 64)
	login_info := &lib.ReqGetScore{
		Uid:    ctx.Query("uid"),
		Name:   ctx.Query("name"),
		Credit: credit,
		Score:  score,
	}
	userDB := service.NewUserDB()
	insert := make([]entities.Scores, 0, 0)
	insert = append(insert, entities.Scores{login_info.Uid, utils.GenerateId(), login_info.Name, login_info.Credit, login_info.Score, 0, 0})
	fmt.Println(insert)
	userDB.InsertScoresSql(insert)
	return response.ResponseOperateSuccess()
}

func GetStudentScores(ctx *gin.Context) *response.Response {
	fmt.Printf("begin getting score\n")
	uid := ctx.Query("uid")
	userDB := service.NewUserDB()
	scores, SelectStudentScoreErr := userDB.GetScoresSql(uid)
	if SelectStudentScoreErr != nil {
		log.Println(SelectStudentScoreErr.Error())
		return response.ResponseQueryFailed()
	}
	resScores := lib.ResGetScores{
		Code:  1,
		Data:  scores,
		Count: 0,
	}
	marshal, MarshalErr := json.Marshal(resScores)
	if MarshalErr != nil {
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(marshal)
}
