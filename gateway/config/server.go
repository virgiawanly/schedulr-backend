package config

import (
	"log"

	"github.com/spf13/viper"
	"github.com/virgiawanly/schedulr-backend/shared/logger"
)

type ServerConfig struct {
	Ip      string  `mapstructure:"server_ip"`
	Port    int     `mapstructure:"server_port"`
	Service Service `mapstructure:"service"`
}

type Service struct {
	Authentication string `mapstructure:"authentication"`
	User           string `mapstructure:"user"`
}

func (srv *ServerConfig) LoadConfig() error {
	v := viper.New()
	v.AutomaticEnv()

	envBindings := map[string]string{
		"server_ip":              "SERVER_IP",
		"server_port":            "SERVER_PORT",
		"service.authentication": "SERVICE_AUTHENTICATION_ADDRESS",
		"service.user":           "SERVICE_USER_ADDRESS",
	}

	for key, envVar := range envBindings {
		v.BindEnv(key, envVar)
	}

	v.SetDefault("server_ip", "0.0.0.0")
	v.SetDefault("server_port", 8000)
	v.SetDefault("service.authentication", "localhost:8001")
	v.SetDefault("service.user", "localhost:8002")

	if err := v.Unmarshal(srv); err != nil {
		logger.Errorf("Failed to unmarshal server config: %v", err)
		return err
	}

	log.Printf("Server config: %+v", srv)

	return nil
}
