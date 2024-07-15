package middlewares

import (
	"RestApiBackend/infrastructure"
	"RestApiBackend/internal/features/users"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type TokenValidatorFunction func(app *infrastructure.Application, accessToken string) (bool, *string, error)

type authMiddleware struct {
	userRepository users.UserRepository
	tokenValidator TokenValidatorFunction
}

func NewAuthBearerToken(userRepo users.UserRepository, app *infrastructure.Application, tokenValidator TokenValidatorFunction) gin.HandlerFunc {
	return authMiddleware{userRepository: userRepo, tokenValidator: tokenValidator}.handle(app)
}

func (m authMiddleware) handle(app *infrastructure.Application) gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			context.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			context.AbortWithStatus(http.StatusForbidden)
			return
		}

		tokenString := parts[1]
		token, s, err := m.tokenValidator(app, tokenString)
		if err != nil {
			context.AbortWithStatus(http.StatusForbidden)
			return
		}
		if !token {
			fmt.Printf("Invalid token %v", token)
			context.AbortWithStatus(http.StatusForbidden)
			return
		}
		user, err := m.userRepository.FetchUserById(context.Request.Context(), *s)
		if err != nil {
			context.AbortWithStatus(http.StatusForbidden)
			return
		}
		context.Set("user", &user)
		context.Next()
	}

}
