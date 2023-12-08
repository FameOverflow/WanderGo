package Authentication

import (
	con "SparkForge/Config"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	oss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
)

var accessKeyId = "LTAI5t5rPhn5UbdoLLQncn8r"
var accessKeySecret = "KGJn7jyzCUBbX7wIgCFpVhngQhY7JK"

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
		"user_account": avatar.UserAccount,
	})
}

func STSUpload() {
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	client, err := oss.New("http://oss-cn-hangzhou.aliyuncs.com", accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 填写Bucket名称，例如examplebucket。
	bucketName := "spark-forge"
	// 填写Object的完整路径，完整路径中不能包含Bucket名称，例如exampledir/exampleobject.txt。
	objectName := "test/test.txt"
	// 填写本地文件的完整路径，例如D:\\localpath\\examplefile.txt。
	filepath := "D:\\music.txt"
	bucket, err := client.Bucket(bucketName)
	// 通过STS授权第三方上传文件。
	err = bucket.PutObjectFromFile(objectName, filepath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	fmt.Println("upload success")
}

func STSDownload() {
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	client, err := oss.New("http://oss-cn-hangzhou.aliyuncs.com", accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 填写Bucket名称，例如examplebucket。
	bucketName := "spark-forge"
	// 填写Object的完整路径，完整路径中不能包含Bucket名称，例如exampledir/exampleobject.txt。
	objectName := "test/test.txt"
	// 填写本地文件的完整路径，例如D:\\localpath\\examplefile.txt。
	filepath := "D:\\music.txt"
	bucket, err := client.Bucket(bucketName)
	// 通过STS授权第三方下载文件。
	err = bucket.GetObjectToFile(objectName, filepath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	fmt.Println("download success")
}
