package Api

import (
	oss "SparkForge/OSS"
	au "SparkForge/controller/authentication"
	com "SparkForge/controller/comments"
	ini "SparkForge/controller/init"
	pos "SparkForge/controller/position"
	mid "SparkForge/middleWare"
	util "SparkForge/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Start() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("ruleOfPwd", util.RuleOfPwd)
	}
	//暂未分组
	engine := gin.Default()
	engine.Use(mid.Cors())
	engine.POST("/register/captcha", au.SendEmail)
	engine.POST("/register", au.RegisterHandler)
	engine.POST("/login", au.LoginHandler)
	engine.POST("/exit", au.ExitHandler)
	engine.POST("/names/change", mid.LoginVerification(), au.ChangeNameHandler)
	engine.POST("/passwords/change", mid.LoginVerification(), au.ChangePwdHandler)
	engine.POST("/passwords/forget/captcha", au.ForgotPasswordGetCaptcha)
	engine.POST("/passwords/forget", au.ForgotPassword)
	engine.POST("/avatars/upload", mid.LoginVerification(), au.AvatarUpload)
	engine.POST("/avatars/change", mid.LoginVerification(), au.AvatarChange)
	engine.POST("/avatars/load", au.SendAvatarToFrontend)
	engine.POST("/comments/post", mid.LoginVerification(), com.PostComment)
	engine.POST("/comments/roam", mid.LoginVerification(), pos.Roaming)
	engine.POST("/comments/like", mid.LoginVerification(), com.LikeHandler)
	engine.POST("/test", com.TestComments)
	engine.POST("/places/mark", pos.MarkPlace)
	engine.POST("/sts/get", oss.GetSTS)
	engine.POST("/places/get", mid.LoginVerification(), pos.PositionsHandler)
	engine.POST("/begin/user", ini.LoadPersonalInformation)
	engine.POST("/begin/places", ini.LoadPlacesInformation)
	engine.Run(":8080")
}
