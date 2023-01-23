package lib

type ReqGetScore struct {
	Uid    string  `json:"id"`
	Name   string  `json:"name"`
	Credit float64 `json:"credit"`
	Score  float64 `json:"score"`
}

type ReqGetUser struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}
