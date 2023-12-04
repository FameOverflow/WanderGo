package Authentication

import "github.com/dgrijalva/jwt-go"

type PwdToBeChanged struct {
	OldPwd     string `json:"old_pwd" binding:"required"`
	CurrentPwd string `json:"current_pwd" binding:"required"`
}
type TempUser struct {
	UserAccount  string `json:"user_account" binding:"required"`
	UserPassword string `json:"user_password" binding:"required"`
}
type NameToBeChanged struct {
	UserName string `json:"user_name" binding:"required"`
}
type UserForgottenPre struct {
	UserAccount string `json:"user_account" binding:"required"`
}
type UserForgotten struct {
	UserAccount string `json:"user_account" binding:"required"`
	NewPwd      string `json:"new_pwd" binding:"required"`
	UserCaptcha int    `json:"user_captcha" binding:"required"`
}
type AccctStatus struct {
	Account            string
	TimeOfChangingName int64
}
type MyClaims struct {
	Account string `json:"account"`
	Time    int    `json:"time"`
	jwt.StandardClaims
}

var SigningKey = []byte("lightlightlight")
