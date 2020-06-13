package config

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv - load the env file
func LoadEnv(file string) {
	if err := godotenv.Load(file); err != nil {
		log.Fatal(err)
	}
}