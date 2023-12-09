package Util

import (
	con "SparkForge/configs"
	"log"
)

// 用user_account从数据库中找到user
func GetUser(a string) con.User {
	var u con.User
	err := con.GLOBAL_DB.Model(&con.User{}).Where("user_account = ?", a).First(&u).Error
	if err != nil {
		log.Panicln(err)
		return u
	}
	return u
}
