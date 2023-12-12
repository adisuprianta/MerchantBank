package controllers

import (
	"MerchantBank/app"
	"MerchantBank/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var requestBank struct {
	Amount int64
}

func Topup(c *gin.Context) {
	c.Bind(&requestBank)
}
func CreateAkunBank(c *gin.Context) {
	c.Bind(&requestBank)
	data, exits := c.Get("user")

	if exits {
		if user, ok := data.(models.User); ok {
			bank := models.Bank{Amount: requestBank.Amount, UserId: user.ID}
			result := app.DB.Create(&bank)
			if result.Error != nil {
				c.Status(http.StatusBadRequest)
				return
			}

			c.JSON(http.StatusCreated, gin.H{
				"status":  http.StatusCreated,
				"massege": "successfully create akun bank",
				"data":    bank,
			})
		}
	}

}
