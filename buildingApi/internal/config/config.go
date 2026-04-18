package config

import (
	"os"
	"time"
)

// Configuration management

// What goes here:

// Env loading
// Config structs
// Validation

/*
   Why?
   Central place for app settings
   Easy to test
   Easy to swap env / config files
*/

/*
   You can copy-paste this file into any new project as your starting point. You would just add new fields to the structs and
   Loadfunction as you add new infrastructure (like a Database or AWS S3) to your app.
*/

// Config holds all application configuration
type Config struct {
	Server ServerConfig
	App    AppConfig
}

// ServerConfig contains HTTP server settings
type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

// AppConfig contains application-specific settings
type AppConfig struct {
	Name        string
	Environment string
	Version     string
}

// Load returns application configuration with environment variable support
// Falls back to sensible defaults for local development
func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port:         getEnv("PORT", "4000"),
			ReadTimeout:  getDurationEnv("READ_TIMEOUT", 15*time.Second),
			WriteTimeout: getDurationEnv("WRITE_TIMEOUT", 15*time.Second),
			IdleTimeout:  getDurationEnv("IDLE_TIMEOUT", 60*time.Second),
		},
		App: AppConfig{
			Name:        getEnv("APP_NAME", "Course API"),
			Environment: getEnv("ENVIRONMENT", "development"),
			Version:     getEnv("VERSION", "1.0.0"),
		},
	}
}

// getEnv retrieves environment variable or returns default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getDurationEnv retrieves duration from environment variable or returns default
func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}
