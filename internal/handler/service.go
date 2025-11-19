package handler

import (
	"context"
	"fmt"

	"github.com/arthurhzna/Golang_gRPC/internal/utils"
	"github.com/arthurhzna/Golang_gRPC/pb/service"
)

type IServiceHandler interface {
	// service.HelloWorldServiceServer
	HelloWorld(ctx context.Context, req *service.HelloWorldRequest) (*service.HelloWorldResponse, error)
}

type serviceHandler struct {
	service.UnimplementedHelloWorldServiceServer
}

func (sh *serviceHandler) HelloWorld(ctx context.Context, req *service.HelloWorldRequest) (*service.HelloWorldResponse, error) {

	validationErrors, err := utils.CheckValidation(req)
	if err != nil {
		return nil, err
	}

	if validationErrors != nil {
		return &service.HelloWorldResponse{
			Base: utils.ValidationErrorResponse(validationErrors),
		}, nil
	}

	return &service.HelloWorldResponse{
		Message: fmt.Sprintf("Hello, %s!", req.Name),
		Base:    utils.SuccessResponse("Success"),
	}, nil
}

// func NewServiceHandler() IServiceHandler {
// 	return &serviceHandler{}
// }

func NewServiceHandler() *serviceHandler {
	return &serviceHandler{}
}
