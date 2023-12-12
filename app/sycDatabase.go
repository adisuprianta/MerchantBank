package app

import "MerchantBank/models"

func SycDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Bank{})
	DB.AutoMigrate(&models.ActivityLog{})
}
