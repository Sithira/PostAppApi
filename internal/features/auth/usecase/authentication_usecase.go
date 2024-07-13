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

func (a authenticationUseCase) Register(ctx context.Context, request *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	user, err := a.userRepository.FetchUserByEmail(ctx, request.Email)

	if err != nil {
		return nil, errors.Wrap(err, "authentication.Register.FetchUserByEmail")
	}

	if user != nil {
		return nil, errors.Wrap(err, "authentication.Register.UserExists")
	}

	if request.Password != request.PasswordRetyped {
		return nil, errors.Wrap(err, "authentication.Register.PasswordMismatched")
	}

	hashedPassword, err := utils.HashPassword(request.Password)

	if err != nil {
		return nil, errors.Wrap(err, "authentication.Register.HashPassword")
	}

	newUser, err := a.userRepository.CreateNewUser(ctx, request.Email, request.FirstName, request.LastName, hashedPassword)

	if err != nil {
		return nil, err
	}

	return &dto.RegisterResponse{
		ID: newUser.ID.String(),
	}, nil
}
