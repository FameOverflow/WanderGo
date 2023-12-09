package Util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncryptMd5(pwd string) string {
	pwdByte := []byte(pwd)
	m := md5.New()
	m.Write(pwdByte)
	pwdEncrypt := hex.EncodeToString(m.Sum(nil))
	return pwdEncrypt
}
