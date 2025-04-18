package config

import (
	"os"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	GRPCAddress     string        `env:"GRPC_ADDRESS" envDefault:"localhost:3001"`
	GRPCAddressUser string        `env:"GRPC_ADDRESS_USER" envDefault:"localhost:3002"`
	HTTPPort        string        `env:"HTTP_PORT" envDefault:":3000"`
	IdleTimeout     time.Duration `env:"IDLE_TIMEOUT" envDefault:"30s"`
	ReadTimeout     time.Duration `env:"READ_TIMEOUT" envDefault:"10s"`
	WriteTimeout    time.Duration `env:"WRITE_TIMEOUT" envDefault:"10s"`
	Debug           bool          `env:"DEBUG" envDefault:"false"`
}

// func Load() (*Config, error) {
// 	if os.Getenv("APP_ENV") != "production" {
// 		_ = godotenv.Load()
// 	}

// 	var cfg Config
// 	if err := env.Parse(&cfg); err != nil {
// 		return nil, err
// 	}

// 	return &cfg, nil
// }

func ProvideConfig() (Config, error) {
	if os.Getenv("APP_ENV") != "production" {
		_ = godotenv.Load()
	}
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
