package repository

import (
	"RestApiBackend/internal/features/users/domain"
	"context"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
}
