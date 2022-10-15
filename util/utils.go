package util

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

//Open file .env
func OpenEnv() {
	//Apertura file.
	file, err := os.Open(".env")
	if err != nil {
		log.Fatalf("failed to open")

	}
	err = godotenv.Load(file.Name())
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
