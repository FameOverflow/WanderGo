package Init

import (
	con "SparkForge/configs"
	au "SparkForge/controller/authentication"
	util "SparkForge/util"
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
	user := util.GetUser(acct)
	var u con.User
	con.GLOBAL_DB.Preload("Comments").Where("user_account = ?", user.UserAccount).First(&u)
	//时间排序
	util.NNewComments = u.Comments
	var commentsPayload []au.CommentsPayload
	sort.Sort(util.NNewComments)
	for i := range util.NNewComments {
		commentsPayload = append(commentsPayload, au.CommentsPayload{
			UserAccount: util.NNewComments[i].UserAccount,
			Date:        util.NNewComments[i].Date,
			Text:        util.NNewComments[i].Text,
			PlaceUID:    util.NNewComments[i].PlaceUID,
			CommentUUID: util.NNewComments[i].CommentUUID,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":   "正处于登录状态",
		"comments":  commentsPayload,
		"user_name": user.UserName,
	})
}
func LoadPlacesInformation(ctx *gin.Context) {
	//地点评论
	var places []con.Place
	err := con.GLOBAL_DB.Preload("Comments").Where("is_marked = 1").Find(&places).Error
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
	var commentsPayload []au.CommentsPayload
	for i := range comments {
		commentsPayload = append(commentsPayload, au.CommentsPayload{
			UserAccount: comments[i].UserAccount,
			Date:        comments[i].Date,
			Text:        comments[i].Text,
			PlaceUID:    comments[i].PlaceUID,
			CommentUUID: comments[i].CommentUUID,
		})
	}

	//place_uid是其所在地点的编号
	//comment_uuid是该评论的编号
	ctx.JSON(http.StatusOK, gin.H{
		"comments_in_place": commentsPayload,
	})
}
