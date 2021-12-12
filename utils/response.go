package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	if data == nil || httpStatus != http.StatusOK {
		ctx.JSON(httpStatus, gin.H{
			"code":code,
			"msg":msg,
		})
		return
	}
	ctx.JSON(httpStatus, gin.H{
		"code":code,
		"msg":msg,
		"data": data,
	})
}

func Succes(ctx *gin.Context, data gin.H, msg string){
	Response(ctx, http.StatusOK, 200, data, msg)
}

func Fail(ctx *gin.Context, data gin.H, msg string){
	Response(ctx, http.StatusOK, 400, data, msg)
}
