package dao

import (
	"github.com/xuyunfeng12388/gin_vue/db"
	"github.com/xuyunfeng12388/gin_vue/model"
	"golang.org/x/crypto/bcrypt"
)

type UserDto struct {
	Id uint `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type UserInfo struct {
	User UserDto `json:"user,omitempty"`
	Token string `json:"token,omitempty"`
}

func ToUserDto(user *model.User) UserDto {
	return UserDto{
		Id:	   user.ID,
		Name:  user.Name,
		Phone: user.Phone,
	}
}
func Resiter(name, password, phone string){
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := model.NewUser(name, string(hashPassword), phone)
	db.DB.Create(user)
}

func GetUserByPhone(phone string) (*model.User, error){
	var user model.User
	if err := db.DB.Select([]string{"id", "name", "phone", "created_at", "updated_at"}).Where("phone=?", phone).First(&user).Error; err != nil{
		return nil, err
	}
	return &user, nil
}

func GetUserByUser(UserId uint) (*model.User, error){
	var user model.User
	if err := db.DB.Select([]string{"id","name", "phone", "created_at", "updated_at"}).Where("ID=?", UserId).First(&user).Error; err != nil{
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
