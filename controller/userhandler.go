package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xuyunfeng12388/gin_vue/utils-main"
	"net/http"
)

func Register(ctx *gin.Context){
	// Get the parameter
	//name := ctx.PostForm("name")
	//password := ctx.PostForm("password")
	phone := ctx.PostForm("phone")
	// Data verification
	if len(phone) != 11{
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":422,
			"msg": "the phone num must be 11 digits!",
		})
		return
	}

	// Create user
	// return result
	ctx.JSON(utils.NewSucc("Register success!", gin.H{
		"msg": "Register success",
	}))
}