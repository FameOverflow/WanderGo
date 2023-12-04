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
func GetAccountWithComments(accountID uint) (dbf.User, error) {
	var u dbf.User
	err := dbf.GLOBAL_DB.Preload("Comment").First(&dbf.User{}, accountID).Error
	if err != nil {
		return dbf.User{}, err
	}
	return u, nil
}
func GetPlaceWithComments(placeID uint) (dbf.Place, error) {
	var place dbf.Place
	err := dbf.GLOBAL_DB.Preload("Comment").First(&place, placeID).Error
	if err != nil {
		return dbf.Place{}, err
	}
	return place, nil
}
