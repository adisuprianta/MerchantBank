package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnnectToDB() {
	var err error

	urlDb := fmt.Sprintf(
		"host=localhost user= %s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	DB, err = gorm.Open(postgres.Open(urlDb), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed connect to database")
	}

}
