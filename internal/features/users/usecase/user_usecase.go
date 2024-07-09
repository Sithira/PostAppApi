package usecase

import (
	"RestApiBackend/internal/features/users"
	"RestApiBackend/internal/features/users/dto"
	"RestApiBackend/internal/features/users/entities"
	"RestApiBackend/pkg/http"
	"context"
	"github.com/pkg/errors"
)

type clientUseCase struct {
	userRepository users.UserRepository
}

func NewUserUserCase(repository users.UserRepository) users.UseCase {
	return &clientUseCase{
		userRepository: repository,
	}
}

func (uc clientUseCase) GetUserByEmail(ctx context.Context, email string) (*dto.UserResponse, error) {
	user, err := uc.userRepository.FetchUserByEmail(ctx, email)
	if err != nil {
		return nil, http.NewBadRequest(errors.Wrap(err, ""))
	}
	return toUserResponse(user), nil
}

func toUserResponse(user *entities.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
