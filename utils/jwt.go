package utils

import (
	"github.com/dgrijalva/jwt-go"
)

//var jwtSecretStr = g.VP.GetString("jwt.secret")

type Claims struct {
	UID  int64 `json:"uid"`
	Time int64 `json:"time"`
	jwt.StandardClaims
}

//// 生成token
//func GenerateToken(claims Claims) (string, error) {
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	return token.SignedString([]byte(jwtSecretStr))
//}
//
//// 解析token
//func ParseToken(tokenString string) (*Claims, error) {
//	token, _ := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
//		return jwtSecretStr, nil
//	})
//
//	if token != nil {
//		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
//			return claims, nil
//		}
//	}
//	return nil, errors.New("invalid token")
//}
