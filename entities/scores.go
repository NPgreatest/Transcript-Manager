package entities

type Scores struct {
	Uid      string  `json:"uid"`
	Sid      string  `json:"sid"`
	Name     string  `json:"name"`
	Credit   float64 `json:"credit"`
	Score    float64 `json:"score"`
	Status   int     `json:"status"`
	Classify int     `json:"classify"`
}
