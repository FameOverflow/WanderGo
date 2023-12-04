package Position

import (
	dbf "SparkForge/Database"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 前端传入坐标,返回地区ID
func PositionHandler(ctx *gin.Context) dbf.Address {
	var pos dbf.Address
	err := ctx.ShouldBind(&pos)
	if err != nil {
		log.Println(err)
	}
	var place dbf.Place
	err = dbf.GLOBAL_DB.Model(&dbf.Place{}).Where("top_left_point.x <= ? AND top_left_point.y <= ? AND bottom_right_point.x >= ? AND bottom_right_point.y >= ?", pos.Y, pos.Y, pos.Y, pos.Y).First(&place).Error
	if err != nil {
		log.Println(err)
	}
	return place.CenterPoint
}

// 随机漫游?
func Roaming(ctx *gin.Context) {
	centerPoint := PositionHandler(ctx)
	var places []dbf.Place
	dbf.GLOBAL_DB.Model(&dbf.Place{}).Where("POWER(?-top_left_point.x, 2) + POWER(?-top_left_point.y, 2) <= 160000", centerPoint.X, centerPoint.Y).Find(&place)
	rand.Seed(time.Now().Unix())
	randomIndex := rand.Intn(len(places))
	selectedPlace := places[randomIndex]
	var comments []dbf.Comment
	dbf.GLOBAL_DB.Model(&dbf.Comment{}).Where("place_id = ?", selectedPlace.ID).Find(&comments)
	if len(comments) == 0 {
		return
	}
	randomCommentIndex := rand.Intn(len(comments))
	selectedComment := comments[randomCommentIndex]
	ctx.JSON(http.StatusOK, gin.H{
		"center_point": selectedPlace.CenterPoint,
	})

}
