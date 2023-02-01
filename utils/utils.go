package utils

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claim struct {
	UserId string
	jwt.StandardClaims
}

var jwtKey = []byte("gpa-manager-npgreatest")

func CreateToken(id string, expireDuration time.Duration) (string, error) {
	claims := &Claim{
		UserId: id,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			Issuer:    "mgh",
			Subject:   "User_Token",
			ExpiresAt: time.Now().Add(expireDuration).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func VerifyToken(token string) (string, bool) {
	if token == "" {
		return "", false
	}

	tok, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		fmt.Println("ParseWithClaims error %v", err)
		return "", false
	}

	if claims, ok := tok.Claims.(*Claim); ok && tok.Valid {
		return claims.UserId, true
	} else {
		fmt.Println("%v", err)
		return "", false
	}

}

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
