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
	if phone == "" || password == "" { // Data verification
		utils.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "the phone and password not null!")
		return
	}

	if len(phone) != 11{ // Data verification
		utils.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "the phone num must be 11 digits!")
		return
	}

	if len(password) < 6 { // verification password
		utils.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "Password cannot be less than 6 digits!")
		return
	}

	user, flag := dao.IsPhoneExist(phone)
	if !flag {
		utils.Fail(ctx, nil, "user not exist, please Register!")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil{
		utils.Fail(ctx, nil, "password err!")
		return
	}
	// 生成token
	token, err := common.ReleaseToken(*user)
	if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, 500, nil, "Server Err!")
		log.Printf("token generate error :%v", err)
		return
	}
	userinfo:= dao.UserInfo{
		User  : dao.ToUserDto(user),
		Token : token,
	}
	utils.Succes(ctx, gin.H{"userinfo": userinfo}, "Login success!")
}

func Register(ctx *gin.Context){
	// Get the parameter
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	phone := ctx.PostForm("phone")

	if phone == "" || password == "" { // Data verification
		utils.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "the phone and password not null!")
		return
	}

	if len(phone) != 11{ // Data verification
		utils.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "the phone num must be 11 digits!")
		return
	}

	if len(password) < 6 { // verification password
		utils.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "Password cannot be less than 6 digits!")
		return
	}

	if len(name)== 0 { // if not input the name, we can gice a 10 digit random strring
		name = utils.RandomString(10)
		fmt.Println(name)
	}
	log.Println(name, password, phone)
	if _, flag := dao.IsPhoneExist(phone); flag{
		utils.Response(ctx, http.StatusBadRequest, 400, nil, "User exist!")
		return
	}
	// Create user
	dao.Resiter(name, password, phone)
	// return result
	utils.Succes(ctx, nil, "Register success!")
}

func UserInfo(ctx *gin.Context) {
	phone := ctx.Query("phone")
	if len(phone) != 11{ // Data verification
		utils.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "the phone num must be 11 digits!")
		return
	}
	user, err := dao.GetUserByPhone(phone)
	if err != nil {
		utils.Fail(ctx, nil, fmt.Sprintf("user %s not exiest!", phone))
		return
	}
	utils.Succes(ctx, gin.H{"userinfo": dao.ToUserDto(user)}, "Login success!")
}
