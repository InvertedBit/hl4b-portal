package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL   string
	SessionSecret string
	Port          string
	UploadsDir    string
}

func Load() *Config {
	// Load .env file if it exists
	_ = godotenv.Load()

	cfg := &Config{
		DatabaseURL:   getEnv("DATABASE_URL", "postgres://localhost:5432/hl4b?sslmode=disable"),
		SessionSecret: getEnv("SESSION_SECRET", "change-this-secret"),
		Port:          getEnv("PORT", "3000"),
		UploadsDir:    getEnv("UPLOADS_DIR", "./uploads"),
	}

	if cfg.SessionSecret == "change-this-secret" {
		log.Println("WARNING: Using default session secret. Please set SESSION_SECRET in .env")
	}

	return cfg
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
