package Position

import (
	con "SparkForge/Config"
	"log"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 前端传入坐标,返回地区中心点
func PositionHandler(ctx *gin.Context) con.Address {
	var pos con.Address
	err := ctx.ShouldBind(&pos)
	if err != nil {
		log.Println(err)
	}
	var place con.Place
	err = con.GLOBAL_DB.Model(&con.Place{}).Where("JSON_EXTRACT(top_left_point, '$.x') <= ? AND JSON_EXTRACT(top_left_point, '$.y') >= ? AND JSON_EXTRACT(bottom_right_point, '$.x') >= ? AND JSON_EXTRACT(bottom_right_point, '$.y') <= ?", pos.X, pos.Y, pos.X, pos.Y).First(&place).Error
	if err != nil {
		log.Println(err)
	}
	return place.CenterPoint
}
func GetPos(c con.Address) con.Place {
	var p con.Place
	err := con.GLOBAL_DB.Model(&con.Place{}).Where("JSON_EXTRACT(center_point, '$.x') = ? AND JSON_EXTRACT(center_point, '$.y') = ?", c.X, c.Y).First(&p).Error
	if err != nil {
		log.Println(err)
	}
	return p
}
func PositionHandlerComment(a con.Address) con.Address {
	var place con.Place
	err := con.GLOBAL_DB.Model(&con.Place{}).Where("JSON_EXTRACT(top_left_point, '$.x') <= ? AND JSON_EXTRACT(top_left_point, '$.y') >= ? AND JSON_EXTRACT(bottom_right_point, '$.x') >= ? AND JSON_EXTRACT(bottom_right_point, '$.y') <= ?", a.X, a.Y, a.X, a.Y).First(&place).Error
	if err != nil {
		log.Println(err)
	}
	return place.CenterPoint
}

// 随机漫游?
func Roaming(ctx *gin.Context) {
	centerPoint := PositionHandler(ctx)
	var places []con.Place
	con.GLOBAL_DB.Model(&con.Place{}).Where("POWER(? - JSON_EXTRACT(center_point, '$.x'), 2) + POWER(? - JSON_EXTRACT(center_point, '$.y'), 2) <= 160000", centerPoint.X, centerPoint.Y).Find(&places)
	randomIndex := rand.Intn(len(places))
	selectedPlace := places[randomIndex]
	var pl con.Place
	con.GLOBAL_DB.Preload("Comments").Where("place_uid = ?", selectedPlace.PlaceUID).First(&pl)
	if len(pl.Comments) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "此处没有漫游点",
		})
		return
	}
	randomCommentIndex := rand.Intn(len(pl.Comments))
	selectedComment := pl.Comments[randomCommentIndex]
	ctx.JSON(http.StatusOK, gin.H{
		"place_uid":  selectedPlace.PlaceUID,
		"place_name": selectedPlace.PlaceName,
	})
	ctx.JSON(http.StatusOK, gin.H{
		"text": selectedComment.Text,
	})
	ctx.Data(http.StatusOK, "image/jpeg", selectedComment.PhotoData)
}
