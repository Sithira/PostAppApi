package middlewares

import (
	"RestApiBackend/infrastructure"
	"RestApiBackend/internal/features/users"
	"RestApiBackend/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type authMiddleware struct {
	ur  users.UserRepository
	app *infrastructure.Application
}

func NewAuthBearerToken(userRepo users.UserRepository, app *infrastructure.Application) gin.HandlerFunc {
	return authMiddleware{ur: userRepo}.handle(app)
}

func (am authMiddleware) handle(app *infrastructure.Application) gin.HandlerFunc {
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
		token, s, err := utils.IsValidJwtAccessToken(app, tokenString)
		if err != nil {
			context.AbortWithStatus(http.StatusForbidden)
			return
		}
		if !token {
			context.Status(http.StatusBadRequest)
			context.Abort()
			return
		}
		user, err := am.ur.FetchUserById(context.Request.Context(), *s)
		if err != nil {
			context.AbortWithStatus(http.StatusForbidden)
			return
		}
		context.Set("user", user)
		context.Next()
	}

}
