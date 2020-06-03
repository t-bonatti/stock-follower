package config

import (
	"github.com/apex/log"
	"github.com/caarlos0/env"
)

// Config data
type Config struct {
	DatabaseURL string `env:"DATABASE_URL" envDefault:"postgres://localhost:5432/stock_data?sslmode=disable"`
}

// Get config
func Get() (cfg Config) {
	if err := env.Parse(&cfg); err != nil {
		log.WithError(err).Fatal("failed to load config")
	}
	return
}
