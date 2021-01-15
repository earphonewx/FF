package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
)

type Password struct {
	PlainText  string
	Iterations int
	Salt       string
	CipherText string
}

func (pwd *Password) MkPassword(plainText string) *Password {
	pwd.PlainText = plainText
	pwd.Iterations = 100000
	var err error
	if pwd.Salt, err = GenRandomStr(12); err != nil {
		panic(fmt.Errorf("生成密码失败：%s\n", err))
	}
	dk := pbkdf2.Key([]byte(plainText), []byte(pwd.Salt), pwd.Iterations, 32, sha256.New)
	cipher := base64.StdEncoding.EncodeToString(dk)
	pwd.CipherText = fmt.Sprintf("pbkdf2_sha256$%d$%s$%s", pwd.Iterations, pwd.Salt, cipher)
	return pwd
}
