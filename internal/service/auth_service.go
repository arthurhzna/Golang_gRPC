package service

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/arthurhzna/Golang_gRPC/internal/entity"
	"github.com/arthurhzna/Golang_gRPC/internal/repository"
	"github.com/arthurhzna/Golang_gRPC/internal/utils"
	"github.com/arthurhzna/Golang_gRPC/pb/auth"
	"github.com/google/uuid"
)

type IAuthService interface {
	Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error)
}

type authService struct {
	authRepository repository.IAuthRepository
}

func NewAuthService(authRepository repository.IAuthRepository) IAuthService {
	return &authService{
		authRepository: authRepository,
	}
}

func (as *authService) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {

	if req.Password != req.PasswordConfirmation {
		return &auth.RegisterResponse{
			Base: utils.BadRequestResponse("Password and password confirmation do not match"),
		}, nil
	}

	user, err := as.authRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return &auth.RegisterResponse{
			Base: utils.BadRequestResponse("User already exists"),
		}, nil
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return nil, err
	}

	NewUser := &entity.User{
		Id:        uuid.New().String(),
		FullName:  req.FullName,
		Email:     req.Email,
		Password:  string(hashPassword),
		RoleCode:  entity.UserRoleCustomer,
		CreatedAt: time.Now(),
		CreatedBy: &req.FullName,
	}

	err = as.authRepository.InsertUser(ctx, NewUser)
	if err != nil {
		return nil, err
	}

	return &auth.RegisterResponse{
		Base: utils.SuccessResponse("User registered successfully"),
	}, nil
}
