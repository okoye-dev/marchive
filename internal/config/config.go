package config

import (
	"log"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	Server ServerConfig
	DB     DBConfig
}

type ServerConfig struct {
	Port string `koanf:"port"`
	WriteTimeout int `koanf:"write_timeout"`
	ReadTimeout int `koanf:"read_timeout"`
	IdleTimeout int `koanf:"idle_timeout"`
}

type DBConfig struct {
	DSN string `koanf:"dsn"`
}

func DefaultServerConfig() *ServerConfig {
	return &ServerConfig{
		Port: ":8080",
		WriteTimeout: 15,
		ReadTimeout:  15,
		IdleTimeout:  60,
	}
}

func DefaultDBConfig() *DBConfig {
	return &DBConfig{
		DSN: "postgres://postgres:password@localhost:5432/marchive",
	}
}

func LoadConfig() *Config {
	k := koanf.New(".")

	if err := k.Load(file.Provider("config.yaml"), yaml.Parser()); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	var cfg *Config
	if err := k.Unmarshal("", cfg); err != nil {
		log.Printf("Error unmarshalling: \n%s", err)
	}

	return cfg
}