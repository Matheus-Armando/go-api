package config

import (
    "os"
    "sync"

    "github.com/joho/godotenv"
)

type Config struct {
    Port           string
    JSONServerURL  string
}

var (
    config *Config
    once   sync.Once
)

// GetConfig returns the application configuration
func GetConfig() *Config {
    once.Do(func() {
        // Load .env file if it exists
        godotenv.Load()

        config = &Config{
            Port:          getEnv("PORT", "8080"),
            JSONServerURL: getEnv("JSON_SERVER_URL", "http://localhost:3000"),
        }
    })

    return config
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}