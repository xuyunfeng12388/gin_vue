package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xuyunfeng12388/gin_vue/common"
	"github.com/xuyunfeng12388/gin_vue/dao"
	"github.com/xuyunfeng12388/gin_vue/utils"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context){
		// 获取authorization
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer"){
			utils.Response(ctx, http.StatusUnauthorized, 401, nil, "Unauthorized!")
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			utils.Response(ctx, http.StatusUnauthorized, 401, nil, "Unauthorized!")
			ctx.Abort()
			return
		}
		// 验证通过获取userID
		userId := claims.UserId
		user, err := dao.GetUserByUser(userId)
		if err != nil {
			utils.Response(ctx, http.StatusUnauthorized, 401, nil, "Unauthorized!")
			ctx.Abort()
			return
		}
		// 将user 写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
