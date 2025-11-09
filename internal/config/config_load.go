package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

func Load() (*Config, error) {
	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, fmt.Errorf("loading config error: %w", err)
	}

	return &cfg, nil
}
