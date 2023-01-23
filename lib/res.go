package lib

type ResGetScores struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}
