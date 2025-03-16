package config

import (
	"fmt"
	"time"

	sharedConfig "github.com/virgiawanly/schedulr-backend/shared/config"
)

type ServerConfig struct {
	sharedConfig.ServerConfig `mapstructure:",squash"`
	Service                   Service        `mapstructure:"service"`
	CircuitBreaker            CircuitBreaker `mapstructure:"cb"`
}

type CircuitBreaker struct {
	MaxRequests      uint32        `mapstructure:"half_state_max_requests"`
	ResetInterval    time.Duration `mapstructure:"half_state_reset_interval"`
	OpenStateTimeout time.Duration `mapstructure:"open_state_timeout"`
}

type Service struct {
	User string `mapstructure:"user"`
}

func (srv *ServerConfig) LoadConfig() error {
	defaults := map[string]interface{}{
		"server_ip":                    "0.0.0.0",
		"server_port":                  50051,
		"cb.half_state_max_requests":   5,
		"cb.half_state_reset_interval": "60s",
		"cb.open_state_timeout":        "30s",
		"service.user":                 "localhost:8002",
	}

	envBindings := map[string]string{
		"server_ip":                    "SERVER_IP",
		"server_port":                  "SERVER_PORT",
		"cb.half_state_max_requests":   "CB_HALF_STATE_MAX_REQUESTS",
		"cb.half_state_reset_interval": "CB_HALF_STATE_RESET_INTERVAL",
		"cb.open_state_timeout":        "CB_OPEN_STATE_TIMEOUT",
		"service.user":                 "SERVICE_USER_ADDRESS",
	}

	return sharedConfig.LoadConfig(srv, defaults, envBindings)
}

func (srv *ServerConfig) GetAddress() string {
	return fmt.Sprintf("%s:%d", srv.Ip, srv.Port)
}
