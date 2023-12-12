package repository

import "MerchantBank/models"

func SycDatabase() {
	DB.AutoMigrate(&models.User{})
}
