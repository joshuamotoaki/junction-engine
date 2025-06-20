package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	NEO4J_URI      string `env:"NEO4J_URI"`
	NEO4J_USERNAME string `env:"NEO4J_USERNAME"`
	NEO4J_PASSWORD string `env:"NEO4J_PASSWORD"`
}

func Load() *Config {
	godotenv.Load()

	return &Config{
		NEO4J_URI:      getEnv("NEO4J_URI", "bolt://localhost:7687"),
		NEO4J_USERNAME: getEnv("NEO4J_USERNAME", "neo4j"),
		NEO4J_PASSWORD: getEnv("NEO4J_PASSWORD", "password123"),
	}

}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
