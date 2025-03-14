package config

import (
	"github.com/joho/godotenv"
	"github.com/virgiawanly/schedulr-backend/shared/logger"
)

type Config struct {
	Server *ServerConfig
}

func (c *Config) LoadConfig() error {
	if err := godotenv.Load(); err != nil {
		logger.Info("No .env file found, using system environment variables")
	}

	if c.Server == nil {
		c.Server = &ServerConfig{}
	}

	if err := c.Server.LoadConfig(); err != nil {
		logger.Errorf("Failed to load server configuration: %v", err)
		return err
	}

	return nil
}
