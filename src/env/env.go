package env

import (
	"log"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}
