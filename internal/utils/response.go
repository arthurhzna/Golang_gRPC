package utils

import (
	"github.com/arthurhzna/Golang_gRPC/pb/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func SuccessResponse(message string) *common.BaseResponse {
	return &common.BaseResponse{
		StatusCode: 200,
		Message:    message,
		IsError:    false,
	}
}

func ValidationErrorResponse(validationErrors []*common.ValidateError) *common.BaseResponse {
	return &common.BaseResponse{
		StatusCode:     400,
		Message:        "Validation errors",
		IsError:        true,
		ValidateErrors: validationErrors,
	}
}

func BadRequestResponse(message string) *common.BaseResponse {
	return &common.BaseResponse{
		StatusCode: 400,
		Message:    message,
		IsError:    true,
	}
}

func NotFoundResponse(message string) *common.BaseResponse {
	return &common.BaseResponse{
		StatusCode: 404,
		Message:    message,
		IsError:    true,
	}
}

func UnaunthorizedResponse() error {
	return status.Errorf(codes.Unauthenticated, "Unauthorized")
}
