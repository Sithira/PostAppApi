package users

import (
	"RestApiBackend/internal/features/users/entities"
	"context"
)

const (
	InsertUserQuery = "INSERT INTO users (id, first_name, last_name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, now(), now()) RETURNING id"
)

type UserRepository interface {
	FetchUserByEmail(ctx context.Context, email string) (*entities.User, error)

	FetchUserById(ctx context.Context, userId string) (*entities.User, error)

	CreateNewUser(ctx context.Context, email string, firstName string, lastName string, password string) (*entities.User, error)
}
