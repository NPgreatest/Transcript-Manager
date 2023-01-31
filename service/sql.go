package service

import (
	"awesomeProject/entities"
	"awesomeProject/lib/db"
	"awesomeProject/utils"
	"bufio"
	"fmt"
	"github.com/axgle/mahonia"
	"log"
	"mime/multipart"
	"strconv"
	"strings"
)

func InsertScoresSql(scores []entities.Scores) error {
	sql := BuildInsertScoresSql(scores)
	value := make([]interface{}, 0)
	for i := range scores {
		value = append(value, scores[i].Uid, scores[i].Sid, scores[i].Name, scores[i].Credit, scores[i].Score, scores[i].Status, scores[i].Classify, scores[i].Append)
	}
	prepare, PrepareErr := db.DB.Prepare(sql)
	if PrepareErr != nil {
		return PrepareErr
	}
	_, ExecErr := prepare.Exec(value...)
	if ExecErr != nil {
		return ExecErr
	}
	return nil
}

func SaveScores(files *multipart.FileHeader, name string) error {
	fmt.Println("begin saving scores")
	var enc mahonia.Decoder
	enc = mahonia.NewDecoder("gbk")
	file, err := files.Open()
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(file)
	scores := make([]entities.Scores, 0)
	var count int64 = 0
	for scanner.Scan() {
		s := entities.Scores{}
		content := scanner.Text()
		content = enc.ConvertString(content)
		info := strings.Split(content, " ")
		s.Name = info[0]
		s.Credit, err = strconv.ParseFloat(info[1], 64)
		if err != nil {
			return err
		}
		s.Score, err = strconv.ParseFloat(info[2], 64)
		if err != nil {
			return err
		}
		s.Uid = name
		s.Sid = utils.GenerateId(count)
		if len(info) > 3 && utils.ConverText(info[3]) != -1 {
			s.Status = utils.ConverText(info[3])
		} else {
			s.Status = -1
		}
		if len(info) > 4 && utils.ConverText(info[4]) != -1 {
			s.Classify = utils.ConverText(info[4])
		} else {
			s.Classify = -1
		}
		if len(info) > 5 {
			flt, err := strconv.ParseFloat(info[5], 64)
			if err != nil {
				s.Append = 0
			} else {
				s.Append = flt
			}
		}
		scores = append(scores, s)
		count += 1
	}
	error := InsertScoresSql(scores)
	if err != nil {
		return error
	}
	return nil
}

func GetScoresSql(uid string) ([]entities.Scores, error) {
	sql := "SELECT * FROM scores WHERE uid=?"
	prepare, PrepareErr := db.DB.Prepare(sql)
	if PrepareErr != nil {
		return nil, PrepareErr
	}
	Query, QueryErr := prepare.Query(uid)
	if QueryErr != nil {
		return nil, QueryErr
	}
	scores := make([]entities.Scores, 0)
	for Query.Next() {
		score := entities.Scores{}
		if err := Query.Scan(&score.Uid, &score.Sid, &score.Name, &score.Credit, &score.Score, &score.Status, &score.Classify, &score.Append); err != nil {
			log.Println(err.Error())
		}
		scores = append(scores, score)
	}
	return scores, nil
}

func BuildInsertScoresSql(scores []entities.Scores) string {
	insert := "INSERT INTO scores VALUES"
	buf := strings.Builder{}
	buf.WriteString(insert)
	for index, _ := range scores {
		buf.WriteString("(")
		buf.WriteString("?,?,?,?,?,?,?,?")
		if index == len(scores)-1 {
			buf.WriteString(")")
		} else {
			buf.WriteString("),")
		}
	}
	sql := buf.String()
	return sql
}
