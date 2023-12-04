package MiddleWare

import (
	au "SparkForge/Authentication"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 中间件判断是否处于登录状态
func LoginOrNot() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("Authorization")
		if auth == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "您没有登录",
			})
			ctx.Abort()
		}
		authAll := strings.Split(auth, " ")
		myClaims := &au.MyClaims{}
		token, err := jwt.ParseWithClaims(authAll[1], myClaims, func(token *jwt.Token) (interface{}, error) {
			return []byte(au.SigningKey), nil
		})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "验证登录失败",
			})
			ctx.Abort()
		}
		if !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "验证登录失败",
			})
			ctx.Abort()
		}
		ctx.Next()
	}
}
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
	}
}
