package config

import (
	"os"
	"strconv"
)

// Config holds application configuration.
// Values are read from environment with sensible defaults.
type Config struct {
	ServerPort int
	Env        string // "development" | "production"
}

// Load reads configuration from environment.
func Load() *Config {
	port := 8080
	if p := os.Getenv("PORT"); p != "" {
		if v, err := strconv.Atoi(p); err == nil {
			port = v
		}
	}
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}
	return &Config{
		ServerPort: port,
		Env:        env,
	}
}
