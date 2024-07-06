package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Token string
	Key   string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("не нашел файл .env")
	}

	return &Config{
		Token: getEnv("BOT_TOKEN", ""),
		Key:   getEnv("API_KEY", ""),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
