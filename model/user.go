package model

type User struct {
	ID int `json:"id"`
	Nmae string `json:"nmae"`
	Password string `json:"password"`
	Phone string `json:"phone"`
}
