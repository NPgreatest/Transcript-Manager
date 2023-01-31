package utils

import (
	"github.com/bwmarrin/snowflake"
)

func GenerateId(x int64) string {
	node, err := snowflake.NewNode(x)
	if err != nil {
		return ""
	}
	return node.Generate().String()
}

func ConverText(text string) int {
	switch text {
	case "初修":
		return 0
	case "重修":
		return 1
	case "免修":
		return 2
	case "必修":
		return 0
	case "选修":
		return 1
	case "其他":
		return 2
	case "政治":
		return 2
	}
	return -1
}
