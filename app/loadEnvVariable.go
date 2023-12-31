package app

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVariable() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
