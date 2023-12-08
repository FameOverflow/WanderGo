package Api

import (
	au "SparkForge/Authentication"
	com "SparkForge/Controller/Comments"

	pos "SparkForge/Controller/Position"
	mid "SparkForge/MiddleWare"
	oss "SparkForge/OSS"

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
	engine.POST("/PostComment", mid.LoginVerification(), com.PostComment)
	engine.POST("/Roaming", mid.LoginVerification(), pos.Roaming)
	engine.POST("/Like", mid.LoginVerification(), com.LikeHandler)
	engine.POST("/test", com.TestComments)
	engine.POST("/MarkPlace", pos.MarkPlace)
	engine.POST("/GetSTS", oss.GetSTS)
	engine.POST("/SearchPlaces", mid.LoginVerification(), pos.PositionsHandler)
	engine.POST("Begin", au.LoadPersonalInformation)
	engine.Run(":8080")
}
