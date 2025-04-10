package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	DBUser string
	DBPass string
	DBHost string
	DBName string
	DBPort string
}

func Load() Config {
	if err := godotenv.Load("services/logly/.env"); err != nil {
		log.Println("⚠️ No .env file found, using system env vars")
	}

	cfg := Config{
		Port:   getEnv("PORT", "8080"),
		DBUser: getEnv("DB_USER", "postgres"),
		DBPass: getEnv("DB_PASSWORD", "securepassword"),
		DBHost: getEnv("DB_HOST", "localhost"),
		DBName: getEnv("DB_NAME", "logly"),
		DBPort: getEnv("DB_PORT", "5432"),
	}
	return cfg
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	log.Printf("[WARN] %s not set. Using default: %s", key, fallback)
	return fallback
}
