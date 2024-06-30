package main

import (
	"clientgrpc/internal/config"
	"clientgrpc/internal/grpcclient"
	"log"

	"github.com/MajotraderLucky/ServerGRPC/api/proto/pb"
)

func main() {
	cfg, err := loadConfig("config/config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	runGRPCClient(cfg)
}

func loadConfig(path string) (*config.Config, error) {
	cfg, err := config.LoadConfig(path)
	if err != nil {
		return nil, err // Обработка ошибки перенесена в main
	}
	return cfg, nil
}

func runGRPCClient(cfg *config.Config) {
	conn, err := grpcclient.CreateGRPCConnection(cfg)
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}
	defer conn.Close()

	c := pb.NewSimpleServiceClient(conn)
	response, err := grpcclient.MakeEchoRequest(c, "Hello, server!")
	if err != nil {
		log.Fatalf("could not make echo request: %v", err)
	}

	log.Printf("Greeting: %s", response.GetMessage())
}
