package utils

import (
	"RestApiBackend/infrastructure"
	"RestApiBackend/internal/features/users/entities"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"net/http"
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

// IsValidJwtAccessToken return user-id
func IsValidJwtAccessToken(app *infrastructure.Application, accessToken string) (bool, *string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signin method %v", token.Header["alg"])
		}
		secret := []byte(app.Env.TokenSignerKey)
		return secret, nil
	})
	if err != nil {
		return false, nil, errors.Wrap(err, "jwt_convert_error")
	}

	if !token.Valid {
		return false, nil, errors.Wrap(err, "jwt_convert_error")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		subject := claims["subject"].(string)
		return true, &subject, nil
	}

	return false, nil, nil
}

func GetUserDetailsFromContext(context *gin.Context) (*uuid.UUID, *entities.User) {
	if context.MustGet("user") != nil {
		userIdFromContext := context.MustGet("user").(*entities.User)
		return &userIdFromContext.ID, userIdFromContext
	}
	context.AbortWithStatus(http.StatusForbidden)
	return nil, nil
}
