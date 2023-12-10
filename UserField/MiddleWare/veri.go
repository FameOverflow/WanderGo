package MiddleWare

import (
	util "SparkForge/util"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 中间件判断是否处于登录状态
func LoginVerification() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("Authorization")
		if auth == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "您没有登录",
			})
			ctx.Abort()
		}
		authAll := strings.Split(auth, " ")
		myClaims := &util.MyClaims{}
		token, err := jwt.ParseWithClaims(authAll[1], myClaims, func(token *jwt.Token) (interface{}, error) {
			return []byte(util.SigningKey), nil
		})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "验证登录失败,获取token失败",
			})
			ctx.Abort()
		}
		if !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "验证登录失败,token无效",
			})
			ctx.Abort()
		}
		ctx.Next()
	}
}
