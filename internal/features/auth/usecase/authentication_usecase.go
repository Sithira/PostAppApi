package usecase

import (
	"RestApiBackend/infrastructure"
	"RestApiBackend/internal/features/auth"
	"RestApiBackend/internal/features/auth/dto"
	"RestApiBackend/internal/features/users"
	"RestApiBackend/pkg/utils"
	"context"
)

type authenticationUseCase struct {
	app            infrastructure.Application
	userRepository users.UserRepository
}

func NewAuthenticationUseCase(app infrastructure.Application, repository users.UserRepository) auth.AuthenticationUseCase {
	return &authenticationUseCase{
		app:            app,
		userRepository: repository,
	}
}

func (a authenticationUseCase) Login(ctx context.Context, request *dto.LoginRequest) (*utils.LoginResponse, error) {
	user, err := a.userRepository.FetchUserByEmail(ctx, request.Email)

	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateLoginToken(&a.app, user)

	if err != nil {
		return nil, err
	}

	return &utils.LoginResponse{
		Type:         token.Type,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}

func (a authenticationUseCase) Logout(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
