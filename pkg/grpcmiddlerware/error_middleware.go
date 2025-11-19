package grpcmiddlerware

import (
	"context"
	"log"
	"runtime/debug"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
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
