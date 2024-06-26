package initializers

import (
	"log"
	"github.com/lpernett/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env variales")
	}
}
