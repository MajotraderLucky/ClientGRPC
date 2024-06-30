package grpcclient

import (
	"clientgrpc/internal/config"
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// NewGRPCClientConn создает новое gRPC соединение с заданным адресом.
func NewGRPCClientConn(target string) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, target, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func CreateGRPCConnection(cfg *config.Config) (*grpc.ClientConn, error) {
	creds, err := credentials.NewClientTLSFromFile(cfg.Certs, "")
	if err != nil {
		return nil, err
	}

	return grpc.Dial(cfg.ServerAddress, grpc.WithTransportCredentials(creds))
}
