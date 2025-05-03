package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVar() {
	err := godotenv.Load()

	if err != nil {
		log.Print(err)
		log.Fatal("Error loading .env config file")
	}
}
