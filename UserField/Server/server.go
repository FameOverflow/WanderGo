package Server

import (
	au "SparkForge/Authentication"
	com "SparkForge/Comments"
	mid "SparkForge/MiddleWare"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Start() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("ruleOfPwd", au.RuleOfPwd)
	}
	engine := gin.Default()
	engine.Use(mid.Cors())
	engine.POST("/GetCaptcha", au.SendEmail)
	engine.POST("/Register", au.RegisterHandler)
	engine.POST("/Login", au.LoginHandler)
	engine.POST("/Exit", au.ExitHandler)
	engine.POST("/ChangeName", mid.LoginOrNot(), au.ChangeNameHandler)
	engine.POST("/ChangePwd", mid.LoginOrNot(), au.ChangePwdHandler)
	engine.POST("/ForgetPwdPre", au.ForgotPasswordGetCaptcha)
	engine.POST("/ForgetPwd", au.ForgotPassword)
	engine.POST("/AvatarUpload", mid.LoginOrNot(), au.AvatarUpload)
	engine.POST("AvatarChange", mid.LoginOrNot(), au.AvatarChange)
	engine.POST("AddComment", com.AddComment)
	engine.Run(":8080")
}
