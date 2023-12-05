package Position

import (
	dbf "SparkForge/Database"
	"log"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 前端传入坐标,返回地区中心点
func PositionHandler(ctx *gin.Context) dbf.Address {
	var pos dbf.Address
	err := ctx.ShouldBind(&pos)
	if err != nil {
		log.Println(err)
	}
	var place dbf.Place
	err = dbf.GLOBAL_DB.Model(&dbf.Place{}).Where("JSON_EXTRACT(top_left_point, '$.x') <= ? AND JSON_EXTRACT(top_left_point, '$.y') >= ? AND JSON_EXTRACT(bottom_right_point, '$.x') >= ? AND JSON_EXTRACT(bottom_right_point, '$.y') <= ?", pos.X, pos.Y, pos.X, pos.Y).First(&place).Error
	if err != nil {
		log.Println(err)
	}
	return place.CenterPoint
}
func GetPos(c dbf.Address) dbf.Place {
	var p dbf.Place
	err := dbf.GLOBAL_DB.Model(&dbf.Place{}).Where("JSON_EXTRACT(center_point, '$.x') = ? AND JSON_EXTRACT(center_point, '$.y') = ?", c.X, c.Y).First(&p).Error
	if err != nil {
		log.Println(err)
	}
	return p
}
func PositionHandlerComment(a dbf.Address) dbf.Address {
	var place dbf.Place
	err := dbf.GLOBAL_DB.Model(&dbf.Place{}).Where("JSON_EXTRACT(top_left_point, '$.x') <= ? AND JSON_EXTRACT(top_left_point, '$.y') >= ? AND JSON_EXTRACT(bottom_right_point, '$.x') >= ? AND JSON_EXTRACT(bottom_right_point, '$.y') <= ?", a.X, a.Y, a.X, a.Y).First(&place).Error
	if err != nil {
		log.Println(err)
	}
	return place.CenterPoint
}

// 随机漫游?
func Roaming(ctx *gin.Context) {
	centerPoint := PositionHandler(ctx)
	var places []dbf.Place
	dbf.GLOBAL_DB.Model(&dbf.Place{}).Where("POWER(? - JSON_EXTRACT(top_left_point, '$.x'), 2) + POWER(? - JSON_EXTRACT(top_left_point, '$.y'), 2) <= 160000", centerPoint.X, centerPoint.Y).Find(&places)
	randomIndex := rand.Intn(len(places))
	selectedPlace := places[randomIndex]
	var pl dbf.Place
	dbf.GLOBAL_DB.Preload("Comments").Where("place_uid = ?", selectedPlace.PlaceUID).First(&pl)
	if len(pl.Comments) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "此处没有漫游点",
		})
		return
	}
	randomCommentIndex := rand.Intn(len(pl.Comments))
	selectedComment := pl.Comments[randomCommentIndex]
	ctx.JSON(http.StatusOK, gin.H{
		"center_point": selectedPlace.CenterPoint,
	})
	ctx.JSON(http.StatusOK, gin.H{
		"sentence": selectedComment.Sentence,
	})

}
