package main

import (
	"clientgrpc/internal/config"
	"context"
	"log"
	"os"
	"time"

	pb "github.com/MajotraderLucky/ServerGRPC/api/proto/pb"
	"github.com/dgrijalva/jwt-go"
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

	token, err := generateJWT()
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

func generateJWT() (string, error) {
	jwtSecretKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Токен истекает через 24 часа
		Issuer:    "exampleIssuer",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func init() {
	// Загрузка .env файла при инициализации пакета
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
