package Util

import (
	"log"
	"regexp"
	"strconv"

	"github.com/go-playground/validator/v10"
)

// 验证器验证注册时账号是否符合规范
func RuleOfAc(fl validator.FieldLevel) bool {
	account := fl.Field().Interface().(int)
	reAc := regexp.MustCompile(`^[1-9][0-9]{9}$`)
	if !reAc.MatchString(strconv.Itoa(account)) {
		log.Println()
		return false
	}
	return true
}

// 验证器验证注册时密码是否符合规范
func RuleOfPwd(fl validator.FieldLevel) bool {
	password := fl.Field().Interface().(string)
	rePwd := regexp.MustCompile(`^[a-zA-Z0-9]{8,12}$`)
	if !rePwd.MatchString(password) {
		log.Println()
		return false
	}
	return true
}
