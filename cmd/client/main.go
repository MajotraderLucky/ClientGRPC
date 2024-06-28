package main

import (
	"clientgrpc/internal/grpcclient"
	"context"
	"log"
	"time"

	pb "github.com/MajotraderLucky/ServerGRPC/api/proto/pb"
)

func main() {
	// Использование новой функции для создания gRPC соединения.
	conn, err := grpcclient.NewGRPCClientConn("localhost:50051")
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
