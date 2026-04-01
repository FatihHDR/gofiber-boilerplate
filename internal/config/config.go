package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	AppName string
	AppPort string
	AppEnv  string
}

func Load() AppConfig {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, reading from environment")
	}

	return AppConfig{
		AppName: getEnv("APP_NAME", "myapp"),
		AppPort: getEnv("APP_PORT", "3000"),
		AppEnv:  getEnv("APP_ENV", "development"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
