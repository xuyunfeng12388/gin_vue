package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuyunfeng12388/gin_vue/common"
	"github.com/xuyunfeng12388/gin_vue/dao"
	"github.com/xuyunfeng12388/gin_vue/utils"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Login(ctx *gin.Context){
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")

	if phone == "" && password == "" { // Data verification
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":422,
			"msg": "the phone and password not null!",
		})
		return
	}

	if len(phone) != 11{ // Data verification
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":422,
			"msg": "the phone num must be 11 digits!",
		})
		return
	}

	if len(password) < 6 { // verification password
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":422,
			"msg": "Password cannot be less than 6 digits!",
		})
		return
	}

	user, flag := dao.IsPhoneExist(phone)
	if !flag {
		ctx.JSON(400, gin.H{
			"code":400,
			"msg": "user not exist, please Register!",
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil{
		ctx.JSON(400, gin.H{
			"code":400,
			"msg": "password err!",
		})
		return
	}
	// 生成token
	token, err := common.ReleaseToken(*user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":http.StatusInternalServerError,
			"msg": "Server Err",
		})
		log.Printf("token generate error :%v", err)
		return
	}
	userinfo:= dao.UserInfo{
		User  : user,
		Token : token,
	}
	ctx.JSON(200, gin.H{
		"code":200,
		"msg": "Login success",
		"data": userinfo,
	})

}

func Register(ctx *gin.Context){
	// Get the parameter
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	phone := ctx.PostForm("phone")

	if phone == "" && password == "" { // Data verification
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":422,
			"msg": "the phone and password not null!",
		})
		return
	}

	if len(phone) != 11{ // Data verification
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":422,
			"msg": "the phone num must be 11 digits!",
		})
		return
	}

	if len(password) < 6 { // verification password
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":422,
			"msg": "Password cannot be less than 6 digits!",
		})
		return
	}

	if len(name)== 0 { // if not input the name, we can gice a 10 digit random strring
		name = utils.RandomString(10)
		fmt.Println(name)
	}
	log.Println(name, password, phone)

	if _, flag := dao.IsPhoneExist(phone); flag{
		ctx.JSON(400, gin.H{
			"msg": "Iphone exist!",
		})
		return
	}
	// Create user
	dao.Resiter(name, password, phone)
	// return result
	//ctx.JSON(utils.NewSucc("Register success!", gin.H{
	//	"msg": "Register success",
	//}))
	ctx.JSON(200, gin.H{
		"msg": "Register success",
	})
}

func UserInfo(ctx *gin.Context) {
	phone := ctx.Query("phone")
	if len(phone) != 11{ // Data verification
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":422,
			"msg": "the phone num must be 11 digits!",
		})
		return
	}
	user, err := dao.GetUserByPhone(phone)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":400,
			"msg": fmt.Sprintf("phone %s not exiest!", phone),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code":200,
		"msg": "success",
		"data": user,
	})
}
