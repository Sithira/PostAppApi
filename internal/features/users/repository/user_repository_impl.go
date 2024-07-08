package postgres

import (
	"RestApiBackend/internal/features/users/domain"
	"RestApiBackend/internal/features/users/repository"
	"context"
	"database/sql"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r userRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	result := r.db.QueryRow("SELECT * FROM users u WHERE u.email = :email AND u.deleted_at IS NULL")
	if result == nil {

	}
	return nil, nil
}
