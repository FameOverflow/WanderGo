package Comments

import (
	con "SparkForge/configs"
	au "SparkForge/controller/authentication"
	util "SparkForge/util"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LikeHandler(ctx *gin.Context) { //"comment_uuid"
	var star con.Star
	err := ctx.ShouldBind(&star)
	if err != nil {
		log.Println(err)
		return
	}
	star.UserAccount = au.SearchAccount(ctx)
	com := GetComment(star.CommentUUID)
	fmt.Println(com)
	user := util.GetUser(star.UserAccount)
	star.User = user
	star.Comment = com
	err = con.GLOBAL_DB.Model(&con.Star{}).Create(&star).Error
	if err != nil {
		log.Println(err)
		return
	}
	com.StarCnt++
	fmt.Println(com.StarCnt)
	err = con.GLOBAL_DB.Model(&con.Comment{}).Where("id = ?", com.ID).Select("star_cnt").Updates(con.Comment{StarCnt: com.StarCnt}).Error
	if err != nil {
		log.Panicln(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "点赞成功",
	})
}
