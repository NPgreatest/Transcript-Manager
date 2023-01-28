package service

import (
	"awesomeProject/entities"
	"awesomeProject/lib/db"
)

func GetAlgorithm(name string) []entities.Algorithm {
	var res []entities.Algorithm
	var sqlRes []entities.Section
	sql1 := "SELECT * FROM algorithms"
	sql2 := "SELECT low,high,point,comment FROM algorithms WHERE name=?"
	if name == "" {
		db.DB.Select(&sqlRes, sql1)
	} else {
		db.DB.Select(&sqlRes, sql2, name)
	}
	if sqlRes == nil {
		return nil
	}
	flag := sqlRes[0].Name
	var temp []entities.Section
	//fmt.Println(sqlRes)
	for index, value := range sqlRes {
		if value.Name != flag || index == len(sqlRes)-1 {
			flag = value.Name
			app, err := entities.NewAlgorithm(temp[0].Name, temp)
			if err != nil {
				return nil
			}
			res = append(res, *app)
			temp = nil
			temp = append(temp, value)
		} else {
			temp = append(temp, value)
			//fmt.Println(index, temp)
		}
	}
	return res
}
