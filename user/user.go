package user

import (
	"awesomeProject/lib"
	"awesomeProject/lib/db"
	"database/sql"
)

type user struct {
	id   string `json:"id"`
	name string `json:"name"`
}
type userDB struct {
	user
	DB *sql.DB
}

func NewUserDB() *userDB {
	return &userDB{
		DB:   db.DB,
		user: user{},
	}
}

func (s *userDB) InsertScore(uid string, data lib.ReqGetScore) error {
	sql := "INSERT INTO scores(`uid`,`sid`,`name`,`credit`,`score`,`status`,`classify`) VALUES(?,?,?,?,?,?,?)"
	prepare, prepareerr := s.DB.Prepare(sql)
	if prepareerr != nil {
		return prepareerr
	}
	_, ExecErr := prepare.Exec(uid, data.Uid, data.Name, data.Credit, data.Score, 0, 0)
	if ExecErr != nil {
		return ExecErr
	}
	return nil
}
