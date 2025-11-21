package jwt

import (
	"context"
	"strings"

	"github.com/arthurhzna/Golang_gRPC/internal/utils"
	"google.golang.org/grpc/metadata"
)

func ParseTokenFromContext(ctx context.Context) (string, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", utils.UnaunthorizedResponse()
	}

	bearerToken, ok := md["authorization"]
	if !ok {
		return "", utils.UnaunthorizedResponse()
	}

	if len(bearerToken) == 0 {
		return "", utils.UnaunthorizedResponse()
	}

	tokenSplit := strings.Split(bearerToken[0], " ")

	if len(tokenSplit) != 2 {
		return "", utils.UnaunthorizedResponse()
	}

	if tokenSplit[0] != "Bearer" {
		return "", utils.UnaunthorizedResponse()
	}

	return tokenSplit[1], nil
}
