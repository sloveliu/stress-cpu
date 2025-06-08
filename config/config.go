package config

import (
	"log"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type AppConfig struct {
	APIKey string
	Port   string
}

func Load() *AppConfig {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("failed to load API_KEY")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("failed to load PORT")
	}

	if _, err := strconv.Atoi(port); err != nil {
		log.Fatalf("failed to load port: %s, error: %s", port, err.Error())
	}

	return &AppConfig{
		APIKey: apiKey,
		Port:   port,
	}
}
