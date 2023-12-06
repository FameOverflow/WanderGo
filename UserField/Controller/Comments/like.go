package Comment

import (
	au "SparkForge/Authentication"
	con "SparkForge/Config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LikeHandler(ctx *gin.Context) {
	var star con.Star
	err := ctx.ShouldBind(&star)
	if err != nil {
		log.Println(err)
		return
	}
	com := GetComment(star.CommentUID)
	user := au.GetUser(star.UserAccount)
	star.User = user
	star.Comment = com
	err = con.GLOBAL_DB.Model(&con.Star{}).Create(&star).Error
	if err != nil {
		log.Println(err)
		return
	}
	com.StarCnt++
	err = con.GLOBAL_DB.Model(&con.Comment{}).Where("id = ?", com.ID).Select("star_cnt").Updates(com).Error
	if err != nil {
		log.Panicln(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "点赞成功",
	})
}
