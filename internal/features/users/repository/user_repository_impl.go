package repository

import (
	"RestApiBackend/internal/features/users"
	"RestApiBackend/internal/features/users/entities"
	"context"
	"database/sql"
	"github.com/pkg/errors"
)

type userRepository struct {
	db *sql.DB
}

func (r *userRepository) FetchUserById(ctx context.Context, userId string) (*entities.User, error) {
	user := entities.User{}

	statements, err := r.db.PrepareContext(ctx, "SELECT * FROM users u WHERE u.id = $1 AND u.deleted_at IS NULL")

	if err != nil {
		return nil, errors.Wrap(err, "authRepository.FetchUserByEmail.PrepareContext")
	}

	if err := statements.QueryRowContext(ctx, userId).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt); err != nil {
		return nil, errors.Wrap(err, "authRepository.FetchUserByEmail.QueryRowContext")
	}

	return &user, nil
}

func NewUserRepository(db *sql.DB) users.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FetchUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	user := entities.User{}

	statements, err := r.db.PrepareContext(ctx, "SELECT * FROM users u WHERE u.email = $1 AND u.deleted_at IS NULL")

	if err != nil {
		return nil, errors.Wrap(err, "authRepository.FetchUserByEmail.PrepareContext")
	}

	if err := statements.QueryRowContext(ctx, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt); err != nil {
		return nil, errors.Wrap(err, "authRepository.FetchUserByEmail.QueryRowContext")
	}

	return &user, nil
}
