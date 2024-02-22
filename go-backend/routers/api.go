package routers

import (
	"github.com/gin-gonic/gin"
	"go-backend/controllers"
	"go-backend/middleware"
	"go-backend/services"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.POST("/auth/register", controllers.Register)
	router.POST("/auth/login", controllers.Login)
	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/auth/info", controllers.Info)
		authRouter.POST("/auth/logout", controllers.Logout)
	}

}
