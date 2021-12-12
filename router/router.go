package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xuyunfeng12388/gin_vue/controller"
	"github.com/xuyunfeng12388/gin_vue/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//r.Use(middleware.CORSMiddleware(), middleware.RecoverMiddleware())
	//r.Static("/static", "static")
	//r.LoadHTMLGlob("template/*")
	r.GET("/", controller.Ping)

	api := r.Group("/v1")
	{
		api.POST("/register", controller.Register)
		api.POST("/login", controller.Login)

		userinfo := api.Group("/user", middleware.AuthMiddleware())
		{
			userinfo.GET("/info", middleware.AuthMiddleware(),  controller.UserInfo)
		}
	}
	return r
}
