package Util

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	Account string `json:"account"`
	Time    int    `json:"time"`
	jwt.StandardClaims
}

var SigningKey = []byte("lightlightlight")

// 生成token，有效期为2周
func GetToken(ac string) string {
	AcTokenPre := MyClaims{
		ac,
		int(time.Now().Unix()),
		jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 60*60*24*14,
			Issuer:    "Fl0RencEss",
		},
	}
	AcTokenByte := jwt.NewWithClaims(jwt.SigningMethodHS256, AcTokenPre)
	AcToken, err := AcTokenByte.SignedString(SigningKey)
	if err != nil {
		log.Println(err)
		return "错误"
	}
	return AcToken
}
