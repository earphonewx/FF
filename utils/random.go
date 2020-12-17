package utils

import (
	"crypto/rand"
	"errors"
	"math/big"
)

const allowedChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func GenRandomStr(length uint) (str string, err error) {
	max := big.NewInt(int64(len(allowedChars)))
	for i := uint(0); i < length; i++ {
		randNum, e := rand.Int(rand.Reader, max)
		if e != nil {
			return "", errors.New("选取随机数失败,无法生成随机字符串！！！")
		}
		str += string(allowedChars[randNum.Int64()])
	}
	return
}
