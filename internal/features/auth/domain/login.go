package domain

import (
	"context"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SuccessLoginResponse struct {
	Type         string `json:"type"`
	ExpiresAt    string `json:"expires_at"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RegisterUseCase interface {
	RegisterUser(c context.Context)
}
