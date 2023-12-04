package Authentication

import (
	dbf "SparkForge/Database"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	mailer "gopkg.in/gomail.v2"
)

func EncryptMd5(pwd string) string {
	pwdByte := []byte(pwd)
	m := md5.New()
	m.Write(pwdByte)
	pwdEncrypt := hex.EncodeToString(m.Sum(nil))
	return pwdEncrypt
}
func AccountConflictVerification(a string) error { //有错说明不冲突
	var tUser dbf.User
	err := dbf.GLOBAL_DB.Model(&dbf.User{}).Where("user_account = ?", a).First(&tUser).Error
	return err
}
func UserLoginVerification(u dbf.User) (int, error) {
	var tUser dbf.User
	err := dbf.GLOBAL_DB.Model(&dbf.User{}).Where(dbf.User{UserAccount: u.UserAccount}, "user_account").
		First(&tUser).Error
	if err != nil {
		// 输入账号不存在
		return 1, err
	} else {
		// 若账号存在，检测密码是否正确
		err = dbf.GLOBAL_DB.Model(&dbf.User{}).Where(dbf.User{UserAccount: u.UserAccount, UserPassword: EncryptMd5(u.UserPassword)}, "user_account", "user_password").
			First(&tUser).Error
		if err != nil {
			// 密码不正确
			return 2, err
		}
		// 登录成功，没有错误
		return 0, nil
	}
}

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

// 通过token找到对应账号
func SearchAccount(ctx *gin.Context) string {
	auth := ctx.Request.Header.Get("Authorization")
	authAll := strings.Split(auth, " ")
	myClaims := &MyClaims{}
	_, err := jwt.ParseWithClaims(authAll[1], myClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SigningKey), nil
	})
	if err != nil {
		log.Println(err)
		return ""
	}
	return myClaims.Account

}

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

// 改名
func ChangeNameHandler(ctx *gin.Context) {
	var nameToBeChanged NameToBeChanged
	err := ctx.ShouldBind(&nameToBeChanged)
	if err != nil {
		log.Println(err)
		return
	}
	userAcct := SearchAccount(ctx)
	err = dbf.GLOBAL_DB.Model(&dbf.User{}).Where("user_account = ?", userAcct).Select("user_name").Updates(dbf.User{UserName: nameToBeChanged.UserName}).Error
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
	var tempUser dbf.User
	userAcct := SearchAccount(ctx)
	dbf.GLOBAL_DB.Model(&dbf.User{}).Where("user_account = ?").First(&tempUser)
	if tempUser.UserPassword != passwordChanging.OldPwd {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "输入的旧密码错误",
		})
		return
	}
	dbf.GLOBAL_DB.Model(&dbf.User{}).Where("user_account = ?", userAcct).Select("user_password").Updates(dbf.User{UserPassword: EncryptMd5(passwordChanging.CurrentPwd)})
	ctx.JSON(http.StatusOK, gin.H{
		"message": "修改密码成功",
	})
}

// 找回密码
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
	msg.SetHeader("From", "m19870110195@163.com")
	msg.SetHeader("To", u.UserAccount)
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
		dbf.GLOBAL_DB.Model(&dbf.User{}).Where("user_account = ?", u.UserAccount).Select("user_password").Updates(dbf.User{UserPassword: EncryptMd5(u.NewPwd)})
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
