package service

import (
	"awesomeProject/entities"
	"log"
	"strings"
)

func (s *UserDB) InsertScoresSql(scores []entities.Scores) error {
	sql := BuildInsertScoresSql(scores)
	value := make([]interface{}, 0)
	for i := range scores {
		value = append(value, scores[i].Uid, scores[i].Sid, scores[i].Name, scores[i].Credit, scores[i].Score, scores[i].Status, scores[i].Classify)
	}
	prepare, PrepareErr := s.DB.Prepare(sql)
	if PrepareErr != nil {
		return PrepareErr
	}
	_, ExecErr := prepare.Exec(value...)
	if ExecErr != nil {
		return ExecErr
	}
	return nil
}

func (s *UserDB) GetScoresSql(uid string) ([]entities.Scores, error) {
	sql := "SELECT * FROM scores WHERE uid=?"
	prepare, PrepareErr := s.DB.Prepare(sql)
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
		if err := Query.Scan(&score.Uid, &score.Sid, &score.Name, &score.Credit, &score.Score, &score.Status, &score.Classify); err != nil {
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
	for range scores {
		buf.WriteString("(")
		buf.WriteString("?,?,?,?,?,?,?")
		buf.WriteString(")")
	}
	sql := buf.String()
	return sql
}
