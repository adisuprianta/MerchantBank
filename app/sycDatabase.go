package app

import "MerchantBank/models"

func SycDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Transfer{})
	DB.AutoMigrate(&models.ActivityLog{})
}
