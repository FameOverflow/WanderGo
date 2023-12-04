package Comment

import (
	dbf "SparkForge/Database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddComment(ctx *gin.Context) {
	var com dbf.Comment
	err := ctx.ShouldBind(&com)
	if err != nil {
		log.Println(err)
		return
	}
	if com.Sentence != "" || com.PhotoData != nil {
		dbf.GLOBAL_DB.Model(&dbf.Comment{}).Create(&com)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "添加评论成功",
		})
	}
}
func ReadOtherPeopleSComments(ctx *gin.Context) {

}
