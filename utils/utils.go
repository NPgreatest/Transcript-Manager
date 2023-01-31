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
