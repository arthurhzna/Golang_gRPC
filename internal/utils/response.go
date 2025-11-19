package utils

import "github.com/arthurhzna/Golang_gRPC/pb/common"

func SuccessResponse(message string) *common.BaseResponse {
	return &common.BaseResponse{
		StatusCode: 200,
		Message:    message,
		IsError:    false,
	}
}
