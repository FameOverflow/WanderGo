package Authentication

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	con "SparkForge/configs"
	util "SparkForge/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	mailer "gopkg.in/gomail.v2"
)

// 处理前端的注册请求,前端设置一下发送完验证码才能点击注册
func RegisterHandler(ctx *gin.Context) {
	var registerAcct con.User
	err := ctx.ShouldBind(&registerAcct)
	if err != nil { //如果没有填写验证码
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "请检查您的输入,必须包含用户邮箱、用户名、密码以及验证码，密码为8-12位，可包含数字和字母",
		})
		log.Printf("err: %v\n", err)
		return
	}
	registerAcct.UserPassword = util.EncryptMd5(registerAcct.UserPassword)
	if CompareCaptcha(registerAcct.UserCaptcha) {
		con.GLOBAL_DB.Model(&con.User{}).Create(&registerAcct)
		token := util.GetToken(registerAcct.UserAccount)
		ctx.SetCookie("_token", "Bearer "+token, 2592000, "/", "localhost", false, true)
		ctx.Request.Header.Set("Authorization", "Bearer "+token)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "注册成功",
		})
		ctx.String(200, token)
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "验证码错误",
		})
	}

}

// 处理前端的登录请求
func LoginHandler(ctx *gin.Context) {
	var loginAcct con.User
	err := ctx.ShouldBind(&loginAcct)
	if err != nil {
		log.Println(err)
		return
	}
	judgeNum, err := UserLoginVerification(loginAcct)
	if err != nil && judgeNum == 1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "您输入的账号不存在",
		})
		return
	}
	if err != nil && judgeNum == 2 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "您输入的密码错误",
		})
		return
	}
	token := util.GetToken(loginAcct.UserAccount)
	BearerToken := "Bearer " + token
	ctx.SetCookie("_token", BearerToken, 2592000, "/", "localhost", false, true)
	ctx.Request.Header.Set("Authorization", BearerToken)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
	})
	ctx.String(200, token)
}

// 用户退出登录清除cookie和token
func ExitHandler(ctx *gin.Context) {
	ctx.SetCookie("_token", "", 0, "/", "", false, true)
	ctx.Request.Header.Del("Authorization")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "您已成功退出该账号",
	})
}

// 找回密码前先发获取验证码
func ForgotPasswordGetCaptcha(ctx *gin.Context) {
	var u UserForgottenPre
	err := ctx.ShouldBind(&u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "请填写邮箱",
		})
		log.Println(err)
		return
	}
	realEmail := regexp.MustCompile("[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+.[a-zA-Z]{2,}$")
	if !realEmail.MatchString(u.UserAccount) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "请输入正确的邮箱地址",
		})
		return
	}
	msg := mailer.NewMessage()
	TempCaptcha = rand.Intn(900000) + 100000
	randNum := strconv.Itoa(TempCaptcha)
	msg.SetHeader("From", GlobalConfig.Email.UserName)
	msg.SetHeader("To", u.UserAccount)
	msg.SetHeader("Subject", "您的漫GO验证码")
	msg.SetBody("text/html", "<h3>您的漫GO验证码为</h3><p>"+randNum+"<p>")
	dialer := mailer.NewDialer(GlobalConfig.Email.Host, GlobalConfig.Email.Port, GlobalConfig.Email.UserName, GlobalConfig.Email.Password) //这个授权码随便用，刚创的
	if err := dialer.DialAndSend(msg); err != nil {
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "验证码已发送",
	})
}
func ForgotPassword(ctx *gin.Context) {
	var u UserForgotten
	err := ctx.ShouldBind(&u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "未输入新密码",
		})
		log.Println(err)
		return
	}
	if CompareCaptcha(u.UserCaptcha) {
		con.GLOBAL_DB.Model(&con.User{}).Where("user_account = ?", u.UserAccount).Select("user_password").Updates(con.User{UserPassword: util.EncryptMd5(u.NewPwd)})
		ctx.JSON(http.StatusOK, gin.H{
			"message": "您已成功修改密码,请登录",
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "验证码错误",
		})
		fmt.Println(TempCaptcha)
	}

}

// 改名
func ChangeNameHandler(ctx *gin.Context) {
	var nameToBeChanged NameToBeChanged
	err := ctx.ShouldBind(&nameToBeChanged)
	if err != nil {
		log.Println(err)
		return
	}
	userAcct := SearchAccount(ctx)
	err = con.GLOBAL_DB.Model(&con.User{}).Where("user_account = ?", userAcct).Select("user_name").Updates(con.User{UserName: nameToBeChanged.UserName}).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "修改用户名失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "修改用户名成功",
	})
}

// 修改密码
func ChangePwdHandler(ctx *gin.Context) {
	var passwordChanging PwdToBeChanged
	err := ctx.ShouldBind(&passwordChanging)
	if err != nil {
		log.Println(err)
		return
	}
	var tempUser con.User
	userAcct := SearchAccount(ctx)
	con.GLOBAL_DB.Model(&con.User{}).Where("user_account = ?").First(&tempUser)
	if tempUser.UserPassword != passwordChanging.OldPwd {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "输入的旧密码错误",
		})
		return
	}
	con.GLOBAL_DB.Model(&con.User{}).Where("user_account = ?", userAcct).Select("user_password").Updates(con.User{UserPassword: util.EncryptMd5(passwordChanging.CurrentPwd)})
	ctx.JSON(http.StatusOK, gin.H{
		"message": "修改密码成功",
	})
}

// 通过token找到对应账号
func SearchAccount(ctx *gin.Context) string {
	auth := ctx.Request.Header.Get("Authorization")
	misakura := 0
	if auth == "" {
		authCookie, err := ctx.Request.Cookie("_token")
		if err == nil {
			auth = authCookie.Value
			misakura = 1
		}
	}
	fmt.Println(auth)
	var authAll []string
	if misakura == 0 {
		authAll = strings.Split(auth, " ")
	} else {
		authAll = strings.Split(auth, "+")
	}
	myClaims := &util.MyClaims{}
	_, err := jwt.ParseWithClaims(authAll[1], myClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(util.SigningKey), nil
	})
	if err != nil {
		log.Println(err)
		return ""
	}
	return myClaims.Account

}
