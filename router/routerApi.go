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
		api.POST("/bank", middleware.RequireAuth, controllers.CreateAkunBank)
		trans := api.Group("/transaction", middleware.RequireAuth)
		{
			trans.POST("/topUp", controllers.Topup)
			trans.POST("/transfer", controllers.Transfer)

		}
	}
	r.Run()
}
