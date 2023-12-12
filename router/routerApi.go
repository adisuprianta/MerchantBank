package router

import (
	"MerchantBank/controllers"
	"github.com/gin-gonic/gin"
)

func RouterApi() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.POST("auth/register", controllers.Sigup)
		api.POST("auth/login", controllers.Login)
		//api.GET("auth/logout", controllers.Logout)
	}
	r.Run()
}
