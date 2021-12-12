package dao

import (
	"github.com/xuyunfeng12388/gin_vue/db"
	"github.com/xuyunfeng12388/gin_vue/model"
	"golang.org/x/crypto/bcrypt"
)

type UserInfo struct {
	User *model.User
	Token string
}

func Resiter(name, password, phone string){
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := model.NewUser(name, string(hashPassword), phone)
	db.DB.Create(user)
}

func GetUserByPhone(phone string) (*model.User, error){
	var user model.User
	if err := db.DB.Select([]string{"id","name", "phone"}).Where("phone=?", phone).First(&user).Error; err != nil{
		return nil, err
	}
	return &user, nil
}

func GetUserByUser(UserId uint) (*model.User, error){
	var user model.User
	if err := db.DB.Select([]string{"id","name", "phone"}).Where("ID=?", UserId).First(&user).Error; err != nil{
		return nil, err
	}
	return &user, nil
}

func IsPhoneExist(phone string)(*model.User, bool)  {
	var user model.User
	db.DB.Where("phone=?", phone).First(&user)
	if user.ID != 0 {
		return &user, true
	}
	return  nil, false
}
