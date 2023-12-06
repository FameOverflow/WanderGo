package Position

import (
	con "SparkForge/Config"
	so "SparkForge/Sort"
	"log"
	"math/rand"
	"net/http"
	"sort"

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
	con.GLOBAL_DB.Model(&con.Place{}).Where("(POWER(? - JSON_EXTRACT(center_point, '$.x'), 2) / POWER(0.0035972, 2)) + (POWER(? - JSON_EXTRACT(center_point, '$.y'), 2) / POWER(0.0045, 2)) <= 1", centerPoint.X, centerPoint.Y).Find(&places)
	randomIndex := rand.Intn(len(places))
	selectedPlace := places[randomIndex]
	var pl con.Place
	con.GLOBAL_DB.Preload("Comments").Where("id = ?", selectedPlace.ID).First(&pl)
	if len(pl.Comments) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "此处没有漫游点",
		})
		return
	}
	randomCommentIndex := rand.Intn(4)
	so.HHotComments = pl.Comments
	sort.Sort(so.HHotComments)
	selectedComment := so.HHotComments[randomCommentIndex]
	ctx.JSON(http.StatusOK, gin.H{
		"id":         selectedPlace.ID,
		"place_name": selectedPlace.PlaceName,
	})
	ctx.JSON(http.StatusOK, gin.H{
		"text":      selectedComment.Text,
		"star_cnts": selectedComment.StarCnt,
	})
	ctx.Data(http.StatusOK, "image/jpeg", selectedComment.PhotoData)
}
