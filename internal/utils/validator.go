package utils

import (
	"errors"

	"buf.build/go/protovalidate"
	"github.com/arthurhzna/Golang_gRPC/pb/common"
	"google.golang.org/protobuf/proto"
)

func CheckValidation(req proto.Message) ([]*common.ValidateError, error) {
	if err := protovalidate.Validate(req); err != nil {
		var validationError *protovalidate.ValidationError
		if errors.As(err, &validationError) {
			var validationErrorsResponse []*common.ValidateError = make([]*common.ValidateError, 0) // array that will store the validation errors
			for _, violation := range validationError.Violations {
				validationErrorsResponse = append(validationErrorsResponse, &common.ValidateError{
					Field:   *violation.Proto.Field.Elements[0].FieldName,
					Message: *violation.Proto.Message,
				})
			}
			return validationErrorsResponse, nil
		}
		return nil, err
	}
	return nil, nil
}
