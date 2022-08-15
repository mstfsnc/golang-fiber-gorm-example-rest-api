package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := Load(".env")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
}

func Load(path string) error {
	return godotenv.Load(path)
}

func Get(key string) string {
	return os.Getenv(key)
}
