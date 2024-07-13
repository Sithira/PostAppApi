package users

import (
	"RestApiBackend/internal/features/users/entities"
	"context"
)

type UserRepository interface {
	FetchUserByEmail(ctx context.Context, email string) (*entities.User, error)

	FetchUserById(ctx context.Context, userId string) (*entities.User, error)
}
