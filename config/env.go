package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var BaseURL string

func init() {
	if godotenv.Load() != nil {
		log.Fatal("Error loading .env file")
	}

	BaseURL = os.Getenv("BASE_URL")

}
