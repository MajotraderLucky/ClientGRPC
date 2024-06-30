package main

import (
	"clientgrpc/internal/config"
	"clientgrpc/internal/grpcclient"
	"log"
)

func main() {
	cfg, err := loadConfig("config/config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	clientService := grpcclient.NewGRPCClientService(cfg)
	clientService.RunGRPCClient()
}

func loadConfig(path string) (*config.Config, error) {
	cfg, err := config.LoadConfig(path)
	if err != nil {
		return nil, err // Обработка ошибки перенесена в main
	}
	return cfg, nil
}
