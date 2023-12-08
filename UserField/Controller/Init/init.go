package Init

import (
	au "SparkForge/Authentication"
	con "SparkForge/Config"
	util "SparkForge/Util"
	"log"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

// 载入页面加载个人信息
func LoadPersonalInformation(ctx *gin.Context) {
	//个人评论
	acct := au.SearchAccount(ctx)
	if acct == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "未登录",
		})
		return
	}
	user := au.GetUser(acct)
	var u con.User
	con.GLOBAL_DB.Preload("Comments").Where("id = ?", user.ID).First(&u)
	//时间排序
	util.NNewComments = u.Comments
	sort.Sort(util.NNewComments)
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "已登录",
		"comments": util.NNewComments,
	})
}
func LoadPlacesInformation(ctx *gin.Context) {
	//地点评论
	var places []con.Place
	err := con.GLOBAL_DB.Preload("Comments").Where("is_marked = 0").Find(&places).Error
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "查询地点数据出错",
		})
		return
	}
	for i := range places {
		util.NNewComments = places[i].Comments
		sort.Sort(util.NNewComments)
		places[i].Comments = util.NNewComments
	}
	var comments []con.Comment
	for i := range places {
		comments = append(comments, places[i].Comments...)
	}
	//comment_uid是其所在地点的id
	//comment_uuid是该评论的标识
	ctx.JSON(http.StatusOK, gin.H{
		"comments_in_place": comments,
	})
}
