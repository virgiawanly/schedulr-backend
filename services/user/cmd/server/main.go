package main

import (
	"schedulr-backend/services/user/config"

	"github.com/virgiawanly/schedulr-backend/shared/logger"
)

func main() {
	appConfig := &config.Config{}

	if err := appConfig.LoadConfig(); err != nil {
		logger.Fatalf("Failed to load configuration: %v", err)
	}

	server, err := NewServer(appConfig)
	if err != nil {
		logger.Fatalf("Failed to create server: %v", err)
	}

	server.Run()
}
