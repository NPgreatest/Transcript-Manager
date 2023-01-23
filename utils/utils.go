package utils

import "github.com/bwmarrin/snowflake"

func GenerateId() string {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return ""
	}
	return node.Generate().String()
}
