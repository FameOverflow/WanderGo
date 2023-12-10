package Authentication

import (
	conf "SparkForge/configs"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	mailer "gopkg.in/gomail.v2"
)

var TempCaptcha int

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
	TempCaptcha = rand.Intn(900000) + 100000
	randNum := strconv.Itoa(TempCaptcha)
	msg.SetHeader("From", conf.EMConfig.UserName)
	msg.SetHeader("To", tu.UserAccount)
	msg.SetHeader("Subject", "您的慢漫验证码")
	msg.SetBody("text/html", "<h3>您的慢漫验证码为</h3><p>"+randNum+"<p>")
	dialer := mailer.NewDialer(conf.EMConfig.Host,conf.EMConfig.Port,conf.EMConfig.UserName,conf.EMConfig.Password) //这个授权码随便用，刚创的
	if err := dialer.DialAndSend(msg); err != nil {
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "验证码已发送",
	})
}
func CompareCaptcha(c int) bool {
	if c == TempCaptcha {
		return true
	} else {
		return false
	}
}
