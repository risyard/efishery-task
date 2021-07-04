package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	PORT          string
	Secret        string
	Key           string
	CacheDuration int
)

func InitConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error getting environment variables for config")
	}

	PORT = os.Getenv("SERVER_PORT")
	Secret = os.Getenv("SECRET")
	Key = os.Getenv("KEY")
	CacheDuration, _ = strconv.Atoi(os.Getenv("CACHE_DURATION"))
}
