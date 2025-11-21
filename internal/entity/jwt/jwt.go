package jwt

import (
	"context"
	"fmt"
	"os"

	"github.com/arthurhzna/Golang_gRPC/internal/utils"
	"github.com/golang-jwt/jwt/v5"
)

type jwtEntityContextKey string

var jwtEntityContextKeyValue jwtEntityContextKey = "jwtEntity"

type JwtClaims struct {
	jwt.RegisteredClaims
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
}

func (jc *JwtClaims) SetToContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, jwtEntityContextKeyValue, jc)
}

func GetClaimsFromToken(jwtToken string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, utils.UnaunthorizedResponse()
	}

	if !token.Valid {
		return nil, utils.UnaunthorizedResponse()
	}

	if claims, ok := token.Claims.(*JwtClaims); ok {
		return claims, nil
	}

	return nil, utils.UnaunthorizedResponse()
}
