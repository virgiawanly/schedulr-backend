package config

import (
	"github.com/joho/godotenv"
	"github.com/virgiawanly/schedulr-backend/shared/logger"
)

type Config struct {
	Server *ServerConfig
}

func (c *Config) LoadConfig() error {
	LoadEnv()
	c.Server = &ServerConfig{}

	configs := []struct {
		Name   string
		Loader func() error
	}{
		{"server", c.Server.LoadConfig},
	}

	for _, cfg := range configs {
		if err := cfg.Loader(); err != nil {
			logger.Errorf("Failed to load %s configuration: %v", cfg.Name, err)
			return err
		}
	}

	return nil
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		logger.Info("No .env file found, using system environment variables")
	}
}
