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
var requestTopup struct {
	Amount int64
	BankId uint
}

func Topup(c *gin.Context) {
	c.Bind(&requestTopup)
	data, exits := c.Get("user")

	if exits {
		if user, ok := data.(models.User); ok {
			var bank models.Bank

			app.DB.First(&bank, "user_id=? AND id=?", user.ID, requestTopup.BankId)

			if bank.ID == 0 {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "data not found",
				})
				return
			}
			totoalAmount := bank.Amount + requestTopup.Amount
			app.DB.Model(&bank).Updates(models.Bank{
				Amount: totoalAmount,
			})
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"massege": "successfully top up amount",
				"data":    bank,
			})
		}
	}
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
