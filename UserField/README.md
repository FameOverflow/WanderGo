# WanderGo
宇宙超级无敌暴龙战士小分队的校园地图项目——漫Go
## 以下皆经POSTMAN测试成功
## 暂未分组
	engine := gin.Default()
	engine.Use(mid.Cors())
### 发送邮箱验证码 "user_account"(string)及"user_password(string)"
 	engine.POST("/GetCaptcha", au.SendEmail)
### 注册 "user_account"及"user_name"(string)及"user_password"及"user_captcha"(int)(如果产品需要用户名)
 	engine.POST("/Register", au.RegisterHandler)
### 登录 "user_account"及"user_password"
 	engine.POST("/Login", au.LoginHandler)
### 退出登录 无需请求体
 	engine.POST("/Exit", au.ExitHandler)
### 改名 "user_name"
	engine.POST("/ChangeName", mid.LoginOrNot(), au.ChangeNameHandler)
### 改密 "old_pwd"及"current_pwd"
	engine.POST("/ChangePwd", mid.LoginOrNot(), au.ChangePwdHandler)
### 忘记密码，发送邮箱验证码 "user_account"
 	engine.POST("/ForgetPwdPre", au.ForgotPasswordGetCaptcha)
### 忘记密码，修改为新密码 "new_pwd"及"user_account"及"user_captcha"
	engine.POST("/ForgetPwd", au.ForgotPassword)
### 上传头像 key为"image"
 	engine.POST("/AvatarUpload", mid.LoginOrNot(), au.AvatarUpload)
### 修改头像 key为"image"
 	engine.POST("AvatarChange", mid.LoginOrNot(), au.AvatarChange)
### 发布评论 (注：图和文字至少传一个)"x"及"y"
 	engine.POST("PostComment",  mid.LoginOrNot(),com.AddComment)
### 随机漫游 "x","y"
  	engine.POST("/Roaming", mid.LoginVerification(), pos.Roaming)
	engine.Run(":8080")