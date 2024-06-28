package main

import (
	"clientgrpc/internal/config"
	"clientgrpc/internal/grpcclient"
	"context"
	"log"
	"time"

	pb "github.com/MajotraderLucky/ServerGRPC/api/proto/pb"
)

func main() {
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	conn, err := grpcclient.NewGRPCClientConn(cfg.ServerAddress)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewSimpleServiceClient(conn)

	reqCtx, reqCancel := context.WithTimeout(context.Background(), time.Second)
	defer reqCancel()

	r, err := c.Echo(reqCtx, &pb.EchoRequest{Message: "Hello, server!"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
