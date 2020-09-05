package router

import (
	"gin-user/controller"
	"gin-user/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	userGroup := r.Group("/user")
	{
		userController := controller.NewUserController()
		userGroup.POST("/register", userController.Create)
		userGroup.POST("/login", userController.Login)
		userGroup.GET("/info", middleware.AuthMiddleware(), userController.Show)
	}
	return r
}
