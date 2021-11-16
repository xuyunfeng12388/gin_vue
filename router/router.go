package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xuyunfeng12388/gin_vue/controller"
)

func Run(){
	r := gin.Default()
	// r.GET("/ping", controller.Ping)

	r.GET("/api/user/register", controller.Register)
	r.Run(":8080")
}
