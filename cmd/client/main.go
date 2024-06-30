package main

import (
	"clientgrpc/internal/config"
	"clientgrpc/internal/grpcclient"
	"log"
)

func main() {
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	clientService := grpcclient.NewGRPCClientService(cfg)
	clientService.RunGRPCClient()
}
