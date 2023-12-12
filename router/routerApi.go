package router

import (
	"MerchantBank/controllers"
	"MerchantBank/middleware"
	"github.com/gin-gonic/gin"
)

func RouterApi() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.POST("auth/register", controllers.Sigup)
		api.POST("auth/login", controllers.Login)
		api.GET("logout", middleware.RequireAuth, controllers.Logout)
		bank := r.Group("/bank", middleware.RequireAuth)
		{
			bank.POST("create", controllers.CreateAkunBank)
			bank.POST("topUp", controllers.Topup)
		}
	}
	r.Run()
}
