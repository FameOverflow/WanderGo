package Comment

import (
	au "SparkForge/Authentication"
	con "SparkForge/Config"
	pos "SparkForge/Controller/Position"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func PostComment(ctx *gin.Context) {
	var com con.Comment
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
	data := time.Now().Format("2006-01-02 15:04:05")
	com.Data = data
	currentTime := strconv.FormatInt(time.Now().Unix(), 10)
	com.CommentUID = au.EncryptMd5(user.UserAccount + currentTime)
	if com.Sentence != "" || com.PhotoData != nil {
		err := con.GLOBAL_DB.Model(&con.Comment{}).Create(&com).Error
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
func GetAccountWithComments(accountID uint) (con.User, error) {
	var u con.User
	err := con.GLOBAL_DB.Preload("Comment").First(&con.User{}, accountID).Error
	if err != nil {
		return con.User{}, err
	}
	return u, nil
}
func GetPlaceWithComments(placeID uint) (con.Place, error) {
	var place con.Place
	err := con.GLOBAL_DB.Preload("Comment").First(&place, placeID).Error
	if err != nil {
		return con.Place{}, err
	}
	return place, nil
}
func TestComments(ctx *gin.Context) {
	// var com dbf.User
	// dbf.GLOBAL_DB.Preload("Comments").Take(&com)
	// fmt.Println(com)
	var p con.Place
	con.GLOBAL_DB.Preload("Comments").Take(&p)
	fmt.Println(p)
}
func GetComment(c string) con.Comment {
	var com con.Comment
	err := con.GLOBAL_DB.Model(&con.Comment{}).Where("comment_uid = ?", c).First(&com)
	if err != nil {
		log.Println(err)
		return con.Comment{}
	}
	return com
}

func HandleNewComments(ctx *gin.Context) {
	err := con.GLOBAL_DB.Model(&con.Comment{}).Find(&newComments).Error
	if err != nil {
		log.Println(err)
		return
	}
	sort.Sort(newComments)
	ctx.JSON(http.StatusOK, newComments)
}

func HandleHotComments(ctx *gin.Context) {
	err := con.GLOBAL_DB.Model(&con.Comment{}).Find(&hotComments).Error
	if err != nil {
		log.Println(err)
		return
	}
	sort.Sort(hotComments)
	ctx.JSON(http.StatusOK, hotComments)
}
