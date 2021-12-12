package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xuyunfeng12388/gin_vue/common"
	"github.com/xuyunfeng12388/gin_vue/dao"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context){
		// 获取authorization
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer"){
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":http.StatusUnauthorized,
				"mes": "权限不够",
			})
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":http.StatusUnauthorized,
				"mes": "权限不够2",
			})
			ctx.Abort()
			return
		}
		// 验证通过获取userID
		userId := claims.UserId
		user, err := dao.GetUserByUser(userId)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":http.StatusUnauthorized,
				"mes": "权限不够3",
			})
			ctx.Abort()
			return
		}
		// 将user 写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
