package main

import (
	"context"
	"log"
	"net"
	"os"
	"runtime/debug"

	"github.com/arthurhzna/Golang_gRPC/internal/handler"
	"github.com/arthurhzna/Golang_gRPC/pb/service"
	"github.com/arthurhzna/Golang_gRPC/pkg/database"
	"github.com/arthurhzna/Golang_gRPC/pkg/grpcmiddlerware"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func ErrorMiddleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
			debug.PrintStack() // print the stack trace or panic stack trace

			err = status.Errorf(codes.Internal, "Internal server error: %v", r)
		}

	}()

	res, err := handler(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}
	return res, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// db := database.ConnectDb(context.Background(), os.Getenv("DB_URL"))
	database.ConnectDb(context.Background(), os.Getenv("DB_URL"))

	lis, err := net.Listen("tcp", ":50051")
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

	serviceHandler := handler.NewServiceHandler()

	grpcServer.Serve(lis)

	service.RegisterHelloWorldServiceServer(grpcServer, serviceHandler)

}
