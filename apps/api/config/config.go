package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	BaseURL        string
	DatabaseURL    string
	AllowedOrigins []string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		Port:        getEnv("PORT", "4006"),
		BaseURL:     getEnv("BASE_URL", ""),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}

	if cfg.BaseURL == "" {
		return nil, fmt.Errorf("BASE_URL is required")
	}

	if cfg.DatabaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required")
	}

	if raw := os.Getenv("ALLOWED_ORIGINS"); raw != "" {
		cfg.AllowedOrigins = strings.Split(raw, ",")
	} else {
		cfg.AllowedOrigins = []string{"http://localhost:3000"}
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
