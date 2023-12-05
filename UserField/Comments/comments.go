package Comment

import (
	au "SparkForge/Authentication"
	dbf "SparkForge/Database"
	pos "SparkForge/Position"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddComment(ctx *gin.Context) {
	var com dbf.Comment
	err := ctx.ShouldBind(&com)
	if err != nil {
		log.Println(err)
		return
	}
	centerPoint := pos.PositionHandlerComment(com.Position)
	place := pos.GetPos(centerPoint)
	com.UserAccount = au.SearchAccount(ctx)
	user := au.GetUser(com.UserAccount)
	com.User = user
	com.Place = place
	currentTime := strconv.FormatInt(com.CreatedAt.Unix(), 10)
	com.CommentUID = au.EncryptMd5(user.UserAccount + currentTime)
	if com.Sentence != "" || com.PhotoData != nil {
		err := dbf.GLOBAL_DB.Model(&dbf.Comment{}).Create(&com).Error
		if err != nil {
			log.Println(err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message":   "添加评论成功",
			"离你最近的中心点为": centerPoint,
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
func TestComments(ctx *gin.Context) {
	// var com dbf.User
	// dbf.GLOBAL_DB.Preload("Comments").Take(&com)
	// fmt.Println(com)
	var p dbf.Place
	dbf.GLOBAL_DB.Preload("Comments").Take(&p)
	fmt.Println(p)
}
