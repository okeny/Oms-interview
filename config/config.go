package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

// Config holds the database connection parameters
type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	SSLMode  string
}

func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	return nil
}

// LoadConfig loads DB configuration from environment or config provider
func LoadConfig() (*Config, error) {
	var missing []string

	get := func(key string) string {
		val, err := Get(key)
		if err != nil {
			missing = append(missing, key)
		}
		return val
	}

	cfg := &Config{
		User:     get("DB_USER"),
		Password: get("DB_PASSWORD"),
		Host:     get("DB_HOST"),
		Port:     get("DB_PORT"),
		Name:     get("DB_NAME"),
		SSLMode:  get("DB_SSLMODE"),
	}

	if len(missing) > 0 {
		return nil, fmt.Errorf("missing required config values: %v", missing)
	}

	return cfg, nil
}
