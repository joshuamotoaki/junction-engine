package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	// Server
	Debug  bool   `env:"DEBUG"`
	Env    string `env:"ENV"` // e.g., development, production
	Port   string `env:"PORT"`
	AppURL string `env:"APP_URL"`

	// Database
	Neo4jURI      string `env:"NEO4J_URI"`
	Neo4jUsername string `env:"NEO4J_USERNAME"`
	Neo4jPassword string ` env:"NEO4J_PASSWORD"`

	// Authentication
	CASServerURL string        `env:"CAS_SERVER_URL"`
	JWTSecret    string        `env:"JWT_SECRET"`
	JWTExpiry    time.Duration `env:"JWT_EXPIRY"`
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	jwtExpiryHours, _ := strconv.Atoi(getEnv("JWT_EXPIRY_HOURS", "24"))

	return &Config{
		Debug:  getEnv("DEBUG", "false") == "true",
		Env:    getEnv("ENV", "dev"),
		Port:   getEnv("PORT", "8080"),
		AppURL: getEnv("APP_URL", "http://localhost:8080"),

		Neo4jURI:      getEnv("NEO4J_URI", "bolt://localhost:7687"),
		Neo4jUsername: getEnv("NEO4J_USERNAME", "neo4j"),
		Neo4jPassword: getEnv("NEO4J_PASSWORD", "password123"),

		CASServerURL: getEnv("CAS_SERVER_URL", "https://fed.princeton.edu/cas"),
		JWTSecret:    getEnv("JWT_SECRET", "your-secret-key-change-this-in-production"),
		JWTExpiry:    time.Duration(jwtExpiryHours) * time.Hour,
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
