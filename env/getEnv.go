package env

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

var KEY string

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	KEY = os.Getenv("KEY")
}
