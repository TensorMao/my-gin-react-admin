package routers

import (
	"github.com/gin-gonic/gin"
	"go-backend/controllers"
)

func SetupBaseRouter(Router *gin.RouterGroup) {

	api := Router.Group("/base")
	{
		//api.POST("/login", controllers.LoginHandler)
		api.GET("/captcha", controllers.GetCaptcha)
	}

}
