package repository

import (
	"RestApiBackend/internal/features/users"
	"RestApiBackend/internal/features/users/entities"
	"context"
	"database/sql"
)

type userRepository struct {
	db *sql.DB
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
		return nil, err
	}

	if err := statements.QueryRowContext(ctx, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt); err != nil {
		return nil, err
	}

	return &user, nil
}
