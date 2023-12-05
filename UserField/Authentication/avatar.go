package Authentication

import (
	con "SparkForge/Config"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AvatarUpload(ctx *gin.Context) {
	var ava con.Avatar // 前端传图
	err := ctx.ShouldBind(&ava)
	if err != nil {
		log.Println(err)
		return
	}
	ava.UserAccount = SearchAccount(ctx)
	file, _, err := ctx.Request.FormFile("image") // 头像的 key 叫 "image"
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "上传头像失败"})
		log.Println(err)
		return
	}
	defer file.Close()
	// 读取文件数据
	buffer, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
		return
	}
	ava.AvatarData = buffer
	con.GLOBAL_DB.Model(&con.Avatar{}).Create(&ava)
	ctx.JSON(http.StatusOK, gin.H{"message": "你有头像了！"})
}
func AvatarChange(ctx *gin.Context) {
	var ava con.Avatar // 前端传图
	err := ctx.ShouldBind(&ava)
	if err != nil {
		log.Println(err)
		return
	}
	ava.UserAccount = SearchAccount(ctx)
	file, _, err := ctx.Request.FormFile("image")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	buffer, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
		return
	}
	con.GLOBAL_DB.Model(&con.Avatar{}).Where("user_name = ?", ava.UserAccount).Select("image_data").Updates(con.Avatar{AvatarData: buffer})
	ctx.JSON(http.StatusOK, gin.H{"message": "换上新头像了！"})
}

func SendAvatarToFrontend(ctx *gin.Context) {
	var avatar con.Avatar
	userAccount := SearchAccount(ctx)
	err := con.GLOBAL_DB.Where("user_account = ?", userAccount).First(&avatar).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "未找到用户头像"})
		return
	}

	// 将头像数据发送给前端
	ctx.Data(http.StatusOK, "image/jpeg", avatar.AvatarData)
	ctx.JSON(http.StatusOK, gin.H{
		"avatar_id": avatar.AvatarID,
	})
}
