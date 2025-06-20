package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl string
}

func Load() *Config {
	godotenv.Load()

	return &Config{
		DBUrl: getEnv("DB_URL", ""),
	}

}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}