package config

import (
	"github.com/caarlos0/env/v6"
)

// Config stuct to store the application configuration
type Config struct {
	DatabaseHost     string `env: "DATABASE_HOST"`
	DatabasePort     string `env: "DATBASE_PORT"`
	DatabaseName     string `env: "DATABASE_NAME"`
	DatabaseUser     string `env: "DATABASE_USER"`
	DatabasePassword string `env: "DATABASE_PASSWORD"`
}

func (c *Config) ParseEnv() *Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
	}

	return &cfg
}
