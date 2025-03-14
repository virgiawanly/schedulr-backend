package config

import sharedConfig "github.com/virgiawanly/schedulr-backend/shared/config"

type ServerConfig struct {
	sharedConfig.ServerConfig `mapstructure:",squash"`
	Service                   Service `mapstructure:"service"`
}

type Service struct {
	Authentication string `mapstructure:"authentication"`
	User           string `mapstructure:"user"`
}

func (srv *ServerConfig) LoadConfig() error {
	defaults := map[string]interface{}{
		"server_ip":              "0.0.0.0",
		"server_port":            8000,
		"service.authentication": "localhost:8001",
		"service.user":           "localhost:8002",
	}

	envBindings := map[string]string{
		"server_ip":              "SERVER_IP",
		"server_port":            "SERVER_PORT",
		"service.authentication": "SERVICE_AUTHENTICATION_ADDRESS",
		"service.user":           "SERVICE_USER_ADDRESS",
	}

	return sharedConfig.LoadConfig(srv, defaults, envBindings)
}
