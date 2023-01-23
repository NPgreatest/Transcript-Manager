package score

import (
	"awesomeProject/entities"
	"awesomeProject/lib"
	"awesomeProject/service"
	"awesomeProject/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func UploadScore(res http.ResponseWriter, req *http.Request) error {
	fmt.Print("begin upload\n")
	credit, cerr := strconv.ParseFloat(req.FormValue("credit"), 64)
	if cerr != nil {
		return cerr
	}
	score, cerr := strconv.ParseFloat(req.FormValue("score"), 64)
	login_info := &lib.ReqGetScore{
		Uid:    req.FormValue("uid"),
		Name:   req.FormValue("name"),
		Credit: credit,
		Score:  score,
	}
	userDB := service.NewUserDB()
	insert := make([]entities.Scores, 0, 0)
	insert = append(insert, entities.Scores{login_info.Uid, utils.GenerateId(), login_info.Name, login_info.Credit, login_info.Score, 0, 0})
	fmt.Println(insert)
	userDB.InsertScoresSql(insert)
	return nil
}

func GetStudentScores(res http.ResponseWriter, req *http.Request) error {
	fmt.Printf("begin getting score\n")
	uid := req.FormValue("uid")
	userDB := service.NewUserDB()
	scores, SelectStudentScoreErr := userDB.GetScoresSql(uid)
	if SelectStudentScoreErr != nil {
		log.Println(SelectStudentScoreErr.Error())
		return SelectStudentScoreErr
	}
	resScores := lib.ResGetScores{
		Code:  1,
		Data:  scores,
		Count: 0,
	}
	marshal, MarshalErr := json.Marshal(resScores)
	if MarshalErr != nil {
		return MarshalErr
	}
	_, err := res.Write(marshal)
	if err != nil {
		return err
	}
	return nil
}
