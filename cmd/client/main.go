package main

import (
	"clientgrpc/internal/config"
	"context"
	"log"
	"time"

	pb "github.com/MajotraderLucky/ServerGRPC/api/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// Загрузка TLS учетных данных
	creds, err := credentials.NewClientTLSFromFile(cfg.Certs, "")
	if err != nil {
		log.Fatalf("could not load tls cert: %v", err)
	}

	// Создание соединения с сервером
	conn, err := grpc.Dial(cfg.ServerAddress, grpc.WithTransportCredentials(creds))
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
