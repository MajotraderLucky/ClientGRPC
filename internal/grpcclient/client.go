package grpcclient

import (
	"context"
	"time"

	"google.golang.org/grpc"
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
