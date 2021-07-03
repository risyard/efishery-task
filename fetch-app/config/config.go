package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT string
	FileName string
	Secret string
)

func InitConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error getting environment variables for config")
	}
	
	PORT = os.Getenv("SERVER_PORT")
	FileName = os.Getenv("FILE_PATH")
	Secret = os.Getenv("SECRET")
}
