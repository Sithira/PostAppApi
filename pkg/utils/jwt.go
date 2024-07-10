package utils

import (
	"RestApiBackend/infrastructure"
	"RestApiBackend/internal/features/users/entities"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type LoginResponse struct {
	Type         string `json:"type"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Claims struct {
	Email string `json:"email"`
	ID    string `json:"id"`
	jwt.MapClaims
}

func GenerateLoginToken(app *infrastructure.Application, user *entities.User) (*LoginResponse, error) {

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject":   user.ID.String(),
		"email":     user.Email,
		"full_name": user.FirstName + " " + user.LastName,
		"exp":       time.Now().Add(time.Minute * 60).Unix(),
		"iat":       time.Now().Unix(),
	})

	tokenString, err := accessToken.SignedString([]byte(app.Env.TokenSignerKey))

	if err != nil {
		return nil, err
	}

	refreshTokenClaims := jwt.MapClaims{
		"type":    "refresh_token",
		"subject": user.ID.String(),
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // Refresh token valid for 7 days
		"iat":     time.Now().Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	refreshTokenString, err := refreshToken.SignedString([]byte(app.Env.TokenSignerKey))
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Type:         "Bearer",
		AccessToken:  tokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
