package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ApiPort   string `envconfig:"API_PORT" required:"true"`
	SecretKey string `envconfig:"SECRET_KEY" required:"true"`
}

func LoadConfig() (*Config, error) {
	var config Config

	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
