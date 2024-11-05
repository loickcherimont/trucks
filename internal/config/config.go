package config

import (
	"log"

	"github.com/joho/godotenv"
)

// Load env file
func LoadVar(fileName string) {

	if err := godotenv.Load(fileName); err != nil {
		log.Fatalf("Error: %q\n", err)
	}
}
