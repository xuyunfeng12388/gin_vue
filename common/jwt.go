package common

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/xuyunfeng12388/gin_vue/model"
	"time"
)

var jwtkey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}


func ReleaseToken(user model.User)(string, error){
	expirationTime := time.Now().Add(7*24*time.Hour)
	claims := &Claims{
		UserId : user.ID,
		StandardClaims:jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		IssuedAt: time.Now().Unix(),
		Issuer: "oceanleran.tech",
		Subject: "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		return "", err
	}
	return tokenString, err
}


func ParseToken(tokenString string)(*jwt.Token, *Claims, error){
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtkey, nil
	})
	return token, claims, err
}