package main

import (
	"github.com/virgiawanly/schedulr-backend/gateway/config"
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

	err = server.Run()
	if err != nil {
		logger.Fatalf("Failed to run server: %v", err)
	}
}
