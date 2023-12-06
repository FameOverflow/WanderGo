package Api

import (
	au "SparkForge/Authentication"
	com "SparkForge/Controller/Comments"
	pos "SparkForge/Controller/Position"
	mid "SparkForge/MiddleWare"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Start() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("ruleOfPwd", au.RuleOfPwd)
	}
	//暂未分组
	engine := gin.Default()
	engine.Use(mid.Cors())
	engine.POST("/GetCaptcha", au.SendEmail)
	engine.POST("/Register", au.RegisterHandler)
	engine.POST("/Login", au.LoginHandler)
	engine.POST("/Exit", au.ExitHandler)
	engine.POST("/ChangeName", mid.LoginVerification(), au.ChangeNameHandler)
	engine.POST("/ChangePwd", mid.LoginVerification(), au.ChangePwdHandler)
	engine.POST("/ForgetPwdPre", au.ForgotPasswordGetCaptcha)
	engine.POST("/ForgetPwd", au.ForgotPassword)
	engine.POST("/AvatarUpload", mid.LoginVerification(), au.AvatarUpload)
	engine.POST("/AvatarChange", mid.LoginVerification(), au.AvatarChange)
	engine.POST("/LoadAvatar", au.SendAvatarToFrontend)
	engine.POST("/AddComment", mid.LoginVerification(), com.PostComment)
	engine.POST("/Roaming", mid.LoginVerification(), pos.Roaming)
	engine.POST("/test", com.TestComments)
	engine.Run(":8080")
}
