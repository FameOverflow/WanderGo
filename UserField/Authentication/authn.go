package Authentication

import (
	"log"
	"net/http"

	con "SparkForge/Config"
	util "SparkForge/Util"

	"github.com/gin-gonic/gin"
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
		ctx.SetCookie("_uuid", util.EncryptMd5(registerAcct.UserAccount), 2592000, "/", "localhost", false, true)
		token := GetToken(registerAcct.UserAccount)
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
	token := GetToken(loginAcct.UserAccount)
	ctx.SetCookie("_token", token, 2592000, "/", "localhost", false, true)
	ctx.Request.Header.Set("Authorization", "Bearer "+token)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
	})
	ctx.String(200, token)
}

// 用户退出登录清除cookie和token
func ExitHandler(ctx *gin.Context) {
	ctx.SetCookie("_uuid", "", 0, "/", "", false, true)
	ctx.Request.Header.Del("Authorization")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "您已成功退出该账号",
	})
}
