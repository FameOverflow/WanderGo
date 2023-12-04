package Authentication

import (
	"log"
	"math/rand"
	"net/http"
	"regexp"
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
	realEmail := regexp.MustCompile("[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+.[a-zA-Z]{2,}$")
	if !realEmail.MatchString(tu.UserAccount) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "请输入正确的邮箱地址",
		})
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
	msg.SetHeader("From", "m19870110195@163.com")
	msg.SetHeader("To", tu.UserAccount)
	msg.SetHeader("Subject", "您的漫GO验证码")
	msg.SetBody("text/html", "<h3>您的漫GO验证码为</h3><p>"+randNum+"<p>")
	dialer := mailer.NewDialer("smtp.163.com", 465, "m19870110195@163.com", "NMNMIOJXWOGJJJJL")
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
