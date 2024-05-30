package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Port        string `env:"PORT" envDefault:":3000"`
	LogLevel    string `env:"LOG_LEVEL" envDefault:"INFO"`
	DatabaseUrl string `env:"DATABASE_URL" envDefault:"database.db"`
}

func New() (*Config, error) {
	_ = godotenv.Load(".env")

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		log.Errorf("Error parsing env vars: %v", err)

		return nil, err
	}
	fmt.Printf("Config: %+v\n", cfg)

	return cfg, nil
}
