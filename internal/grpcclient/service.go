package grpcclient

import (
	"clientgrpc/internal/auth"
	"clientgrpc/internal/config"
	"context"
	"log"
	"time"

	pb "github.com/MajotraderLucky/ServerGRPC/api/proto/pb"
	"google.golang.org/grpc/metadata"
)

func MakeEchoRequest(client pb.SimpleServiceClient, message string) (*pb.EchoResponse, error) {
	token, err := auth.GenerateJWT()
	if err != nil {
		return nil, err
	}

	md := metadata.Pairs("authorization", "Bearer "+token)
	reqCtx, reqCancel := context.WithTimeout(metadata.NewOutgoingContext(context.Background(), md), time.Second)
	defer reqCancel()

	return client.Echo(reqCtx, &pb.EchoRequest{Message: message})
}

type GRPCClientService struct {
	Config *config.Config
}

func NewGRPCClientService(cfg *config.Config) *GRPCClientService {
	return &GRPCClientService{Config: cfg}
}

func (s *GRPCClientService) RunGRPCClient() {
	conn, err := CreateGRPCConnection(s.Config)
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}
	defer conn.Close()

	c := pb.NewSimpleServiceClient(conn)
	response, err := MakeEchoRequest(c, "Hello, server!")
	if err != nil {
		log.Fatalf("could not make echo request: %v", err)
	}

	log.Printf("Greeting: %s", response.GetMessage())
}
