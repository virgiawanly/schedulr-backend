package config

import (
	"github.com/joho/godotenv"
	"github.com/virgiawanly/schedulr-backend/shared/logger"
)

type Config struct {
	DB        *DBConfig
	Server    *ServerConfig
	JWTConfig *JWTConfig
}

func (c *Config) LoadConfig() error {
	LoadEnv()

	c.DB = &DBConfig{}
	c.Server = &ServerConfig{}
	c.JWTConfig = &JWTConfig{}

	configs := []struct {
		Name   string
		Loader func() error
	}{
		{"server", c.Server.LoadConfig},
		{"database", c.DB.LoadConfig},
		{"jwt", c.JWTConfig.LoadConfig},
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
