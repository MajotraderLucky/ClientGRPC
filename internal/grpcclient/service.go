package grpcclient

import (
	"clientgrpc/internal/auth"
	"context"
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
