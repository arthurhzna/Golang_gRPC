package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/arthurhzna/Golang_gRPC/internal/handler"
	"github.com/arthurhzna/Golang_gRPC/internal/repository"
	"github.com/arthurhzna/Golang_gRPC/internal/service"
	"github.com/arthurhzna/Golang_gRPC/pb/auth"
	"github.com/arthurhzna/Golang_gRPC/pkg/database"
	"github.com/arthurhzna/Golang_gRPC/pkg/grpcmiddlerware"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer lis.Close()

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(
		grpcmiddlerware.ErrorMiddleware,
	))

	if os.Getenv("ENVIRONMENT") == "DEV" {
		reflection.Register(grpcServer)
	}

	db := database.ConnectDb(context.Background(), os.Getenv("DB_URL"))
	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository)
	authHandler := handler.NewAuthHandler(authService)

	auth.RegisterAuthServiceServer(grpcServer, authHandler)

	grpcServer.Serve(lis)

}
