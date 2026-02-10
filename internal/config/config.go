package config

import (
	"log"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	Server ServerConfig `koanf:"server"`
	DB     DBConfig     `koanf:"database"`
}

type ServerConfig struct {
	Port            string `koanf:"port"`
	WriteTimeout    int    `koanf:"write_timeout"`
	ReadTimeout     int    `koanf:"read_timeout"`
	IdleTimeout     int    `koanf:"idle_timeout"`
	ShutdownTimeout int    `koanf:"shutdown_timeout"`
}

type DBConfig struct {
	DSN string `koanf:"dsn"`
}

func DefaultServerConfig() *ServerConfig {
	return &ServerConfig{
		Port:            "8080",
		WriteTimeout:    15,
		ReadTimeout:     15,
		IdleTimeout:     60,
		ShutdownTimeout: 10,
	}
}

func DefaultDBConfig() *DBConfig {
	return &DBConfig{
		DSN: "postgres://postgres:password@localhost:5432/marchive",
	}
}

func DefaultConfig() *Config {
	return &Config{
		Server: *DefaultServerConfig(),
		DB:     *DefaultDBConfig(),
	}
}

func LoadConfig() *Config {
	k := koanf.New(".")

	if err := k.Load(file.Provider("config.yaml"), yaml.Parser()); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	cfg := DefaultConfig()
	if err := k.Unmarshal("", cfg); err != nil {
		log.Fatalf("Error unmarshalling: %v", err)
	}

	return cfg
}
