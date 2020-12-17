package utils

import (
	"errors"
	"ff/g"
	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(g.VP.GetString("jwt.secret"))

type Claims struct {
	UID  int64 `json:"uid"`
	Time int64 `json:"time"`
	jwt.StandardClaims
}

// 生成token
func GenerateToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// 解析token
func ParseToken(tokenString string) (*Claims, error) {
	token, _ := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return jwtSecret, nil
	})

	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, errors.New("invalid token")
}
