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
var requestTransfer struct {
	Amount            int64
	BankId            uint
	DestinationBankId int64
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
					"error": "bank account not found",
				})
				return
			}
			totalAmount := bank.Amount + requestTopup.Amount
			app.DB.Model(&bank).Updates(models.Bank{
				Amount: totalAmount,
			})
			logoModel := models.ActivityLog{Action: "top up", UserId: user.ID, Description: "Melakukan update amount"}

			app.DB.Create(&logoModel)
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"massage": "successfully top up amount",
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
			logoModel := models.ActivityLog{Action: "Create", UserId: user.ID, Description: "Melakukan create akun bank"}

			app.DB.Create(&logoModel)

			c.JSON(http.StatusCreated, gin.H{
				"status":  http.StatusCreated,
				"massage": "successfully create account bank",
				"data":    bank,
			})
		}
	}

}

func Transfer(c *gin.Context) {
	c.Bind(&requestTransfer)
	data, exits := c.Get("user")

	if exits {
		if user, ok := data.(models.User); ok {
			var bank models.Bank
			app.DB.First(&bank, "user_id=? AND id=?", user.ID, requestTransfer.BankId)
			var destinationBank models.Bank
			app.DB.First(&destinationBank, requestTransfer.DestinationBankId)

			if bank.ID == 0 {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "bank account not found",
				})
				return
			}
			if destinationBank.ID == 0 {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "destination bank account not found",
				})
				return
			}
			destinationAmount := destinationBank.Amount + requestTransfer.Amount
			totalAmount := bank.Amount - requestTransfer.Amount

			app.DB.Model(&bank).Updates(models.Bank{
				Amount: totalAmount,
			})
			app.DB.Model(&destinationAmount).Updates(models.Bank{
				Amount: destinationAmount,
			})

			logoModel := models.ActivityLog{Action: "transfer", UserId: user.ID, Description: "melakukan transfer dana"}

			app.DB.Create(&logoModel)
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"massage": "successfully transfer amount",
				"data":    bank,
			})
		}
	}
}
