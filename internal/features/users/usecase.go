package users

import (
	"RestApiBackend/internal/features/users/dto"
	"context"
)

type UseCase interface {
	GetUserByEmail(ctx context.Context, email string) (*dto.UserResponse, error)
}
