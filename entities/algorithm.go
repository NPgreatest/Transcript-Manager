package entities

import "fmt"

type Algorithm struct {
	Name     string    `json:"Name"`
	Sections []Section `json:"Sections"`
}

type Section struct {
	Name    string  `json:"Name"`
	Low     float64 `json:"Low"`
	High    float64 `json:"High"`
	Point   float64 `json:"Point"`
	Comment string  `json:"Comment"`
}

func NewAlgorithm(name string, sections []Section) (*Algorithm, error) {
	if len(sections) == 0 {
		return nil, fmt.Errorf("没有内容")
	}
	return &Algorithm{Name: name, Sections: sections}, nil
}

func GetPoint(alg Algorithm, score float64) float64 {
	for _, value := range alg.Sections {
		if score > value.Low && score <= value.High {
			return value.Point
		}
	}
	return 0
}

func GetComment(alg Algorithm, score float64) string {
	for _, value := range alg.Sections {
		if score > value.Low && score <= value.High {
			return value.Comment
		}
	}
	return "ERROR"
}
