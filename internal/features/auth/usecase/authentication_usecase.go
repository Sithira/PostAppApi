package usecase

import (
	"RestApiBackend/infrastructure"
	"RestApiBackend/internal/features/auth"
	"RestApiBackend/internal/features/auth/dto"
	"RestApiBackend/internal/features/users"
	"RestApiBackend/pkg/utils"
	"context"
	"fmt"
	"github.com/pkg/errors"
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

func (a authenticationUseCase) Login(ctx context.Context, request *dto.LoginRequest) (*utils.LoginResponse, *string, error) {

	user, err := a.userRepository.FetchUserByEmail(ctx, request.Email)

	if err != nil {
		return nil, nil, errors.Wrap(err, "authentication.Login.FetchUserByEmail")
	}

	matches := utils.CompareHashAndPassword(request.Password, *user.Password)

	if !matches {
		return nil, nil, errors.Wrap(nil, "authentication.Login.CompareHashAndPassword.Failed")
	}

	token, err := utils.GenerateLoginToken(&a.app, user)
	fmt.Printf("Login token generated for user %s", token)

	if err != nil {
		return nil, nil, errors.Wrap(err, "authentication.Login.GenerateLoginToken")
	}

	userId := user.ID.String()

	return &utils.LoginResponse{
		Type:         token.Type,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, &userId, nil
}

func (a authenticationUseCase) Logout(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
