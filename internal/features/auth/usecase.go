package auth

import (
	"RestApiBackend/internal/features/auth/dto"
	"RestApiBackend/pkg/utils"
	"context"
)

type AuthenticationUseCase interface {
	Login(ctx context.Context, request *dto.LoginRequest) (*utils.LoginResponse, *string, error)

	Logout(ctx context.Context) error

	Register(ctx context.Context, request *dto.RegisterRequest) (*dto.RegisterResponse, error)
}
