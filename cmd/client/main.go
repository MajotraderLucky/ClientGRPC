package main

import (
	"clientgrpc/internal/config"
	"clientgrpc/internal/grpcclient"
	"log"

	pb "github.com/MajotraderLucky/ServerGRPC/api/proto/pb"
)

func main() {
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

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
