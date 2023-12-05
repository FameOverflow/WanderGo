package Comment

import (
	con "SparkForge/Config"
	"log"

	"github.com/gin-gonic/gin"
)

func LikeHandler(ctx *gin.Context) {
	var star con.Star
	err := ctx.ShouldBind(&star)
	if err != nil {
		log.Println(err)
		return
	}
}
