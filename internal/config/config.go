package config

import "time"

type Config struct {
	Port string
	WriteTimeout time.Duration
	ReadTimeout time.Duration
	IdleTimeout time.Duration
}

func DefaultConfig() *Config {
	return &Config{
		Port: "8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}