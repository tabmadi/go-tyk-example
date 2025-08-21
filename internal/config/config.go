// Package config provides configuration management.
package config

import "time"

// Config holds server configuration settings.
type Config struct {
	ListenAddress string
	TimeOut       time.Duration
}

// Load initializes and returns a Config instance with default values.
func Load() *Config {
	cfg := &Config{
		ListenAddress: "localhost:8080",
		TimeOut:       time.Second * 15,
	}

	return cfg
}
