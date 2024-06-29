package main

import (
	"clientgrpc/internal/auth"
	"clientgrpc/internal/config"
	"context"
	"log"
	"time"

	pb "github.com/MajotraderLucky/ServerGRPC/api/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func main() {
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	conn, err := createGRPCConnection(cfg)
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}
	defer conn.Close()

	c := pb.NewSimpleServiceClient(conn)

	response, err := makeEchoRequest(c, "Hello, server!")
	if err != nil {
		log.Fatalf("could not make echo request: %v", err)
	}

	log.Printf("Greeting: %s", response.GetMessage())
}

func createGRPCConnection(cfg *config.Config) (*grpc.ClientConn, error) {
	creds, err := credentials.NewClientTLSFromFile(cfg.Certs, "")
	if err != nil {
		return nil, err
	}

	return grpc.Dial(cfg.ServerAddress, grpc.WithTransportCredentials(creds))
}

func makeEchoRequest(client pb.SimpleServiceClient, message string) (*pb.EchoResponse, error) {
	token, err := auth.GenerateJWT()
	if err != nil {
		return nil, err
	}

	md := metadata.Pairs("authorization", "Bearer "+token)
	reqCtx, reqCancel := context.WithTimeout(metadata.NewOutgoingContext(context.Background(), md), time.Second)
	defer reqCancel()

	return client.Echo(reqCtx, &pb.EchoRequest{Message: message})
}
