package service

import (
	"awesomeProject/entities"
	"fmt"
	"math"
)

func GetGPAGroup(stuname string, algname string) (entities.GPAGroup, error) {
	scores, err := GetScoresSql(stuname)
	group := entities.NewGPAGroup()
	if err != nil {
		return *group, err
	}
	majorgroup := make([]entities.Scores, 0)
	overallgroup := make([]entities.Scores, 0)
	initialgroup := make([]entities.Scores, 0)
	for _, value := range scores {
		switch value.Status {
		case 0: //初修
			initialgroup = append(initialgroup, value)
			overallgroup = append(overallgroup, value)
			if value.Classify == 0 {
				majorgroup = append(majorgroup, value)
			}
		case 1: //重修
			initialgroup = append(initialgroup, value)
			var temp = value
			temp.Score = math.Max(temp.Append, temp.Score)
			overallgroup = append(overallgroup, temp)
			if value.Classify == 0 {
				majorgroup = append(majorgroup, temp)
			}
		default:
			initialgroup = append(initialgroup, value)
			overallgroup = append(overallgroup, value)
			if value.Classify == 0 {
				majorgroup = append(majorgroup, value)
			}
		}
	}
	group.InitialGPA = CalculateGPA(initialgroup, algname)
	group.MajorGPA = CalculateGPA(majorgroup, algname)
	group.OverallGPA = CalculateGPA(overallgroup, algname)
	group.InitialScore = CalculateScore(initialgroup)
	group.MajorScore = CalculateScore(majorgroup)
	group.OverallScore = CalculateScore(overallgroup)
	return *group, nil
}

func CalculateScore(scores []entities.Scores) float64 {
	var res float64 = 0
	var total float64 = 0
	for _, i := range scores {
		res += i.Score * i.Credit
		total += i.Credit
	}
	return res / total
}

func CalculateGPA(scores []entities.Scores, algname string) float64 {
	var res float64 = 0
	var total float64 = 0
	for _, i := range scores {
		if GetAlgorithm(algname) == nil {
			fmt.Println("Something went wrong when calculating GPA")
			return 0
		}
		res += entities.GetPoint(GetAlgorithm(algname)[0], i.Score) * i.Credit
		total += i.Credit
	}
	return res / total
}
