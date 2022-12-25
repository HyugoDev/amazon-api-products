package ConfigApi

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetPORT() string {
	port := ":" + os.Getenv("PORT")

	return port
}

func GetURLMongo() string {
	url := os.Getenv("URL_MONGO")

	return url
}
