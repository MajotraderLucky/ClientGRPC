package main

import (
	"clientgrpc/internal/auth"
	"clientgrpc/internal/config"
	"context"
	"log"
	"time"

	pb "github.com/MajotraderLucky/ServerGRPC/api/proto/pb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func main() {
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	creds, err := credentials.NewClientTLSFromFile(cfg.Certs, "")
	if err != nil {
		log.Fatalf("could not load tls cert: %v", err)
	}

	conn, err := grpc.Dial(cfg.ServerAddress, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewSimpleServiceClient(conn)

	token, err := auth.GenerateJWT()
	if err != nil {
		log.Fatalf("could not generate token: %v", err)
	}

	md := metadata.Pairs("authorization", "Bearer "+token)
	reqCtx, reqCancel := context.WithTimeout(metadata.NewOutgoingContext(context.Background(), md), time.Second)
	defer reqCancel()

	r, err := c.Echo(reqCtx, &pb.EchoRequest{Message: "Hello, server!"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

func init() {
	// Загрузка .env файла при инициализации пакета
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
