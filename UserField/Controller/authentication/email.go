package Authentication

import (
	conf "SparkForge/configs"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	mailer "gopkg.in/gomail.v2"
)

var TempCaptcha UserCaptcha

func SendEmail(ctx *gin.Context) {
	var tu TempUser
	err := ctx.ShouldBind(&tu)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "请填写邮箱和密码",
		})
		log.Println(err)
		return
	}
	err = AccountConflictVerification(tu.UserAccount) //验证该邮箱是否已经注册
	if err == nil {                                   //如果已有数据中存在该邮箱则注册失败
		ctx.JSON(http.StatusConflict, gin.H{
			"message": "该账号已被注册",
		})
		return
	}
	msg := mailer.NewMessage()
	TempCaptcha.Captcha = rand.Intn(900000) + 100000
	TempCaptcha.UserAccount = tu.UserAccount
	TempCaptcha.ExpireTime=time.Now().Add(10 * time.Minute).Unix()
	randNum := strconv.Itoa(TempCaptcha.Captcha)
	msg.SetHeader("From", conf.EMConfig.UserName)
	msg.SetHeader("To", tu.UserAccount)
	msg.SetHeader("Subject", "慢漫注册验证码")
	msg.SetBody("text/html", "<p>您正在注册慢漫，以下是您的验证码，验证码将在十分钟后过期<p><h2>"+randNum+"<h2><p>如果这不是您的邮件请忽略，很抱歉打扰您，请原谅。<p>")
	dialer := mailer.NewDialer(conf.EMConfig.Host,conf.EMConfig.Port,conf.EMConfig.UserName,conf.EMConfig.Password) //这个授权码随便用，刚创的
	if err := dialer.DialAndSend(msg); err != nil {
		log.Println(err)
		return
	}
	conf.GLOBAL_DB.Model(&UserCaptcha{}).Create(&TempCaptcha)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "验证码已发送",
	})
}
func CompareCaptcha(captcha int,Account string) bool {
	var TempCaptcha = UserCaptcha{}
	conf.GLOBAL_DB.Model(&UserCaptcha{}).Where("user_account = ?",Account).First(&TempCaptcha)
	if TempCaptcha.Captcha == captcha && TempCaptcha.ExpireTime > time.Now().Unix(){
		return true
	}
	return false
}
