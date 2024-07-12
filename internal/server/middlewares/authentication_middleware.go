package middlewares

import (
	"RestApiBackend/infrastructure"
	"RestApiBackend/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func BearerTokenMiddleware(app *infrastructure.Application) gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			context.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			context.Abort()
			return
		}

		tokenString := parts[1]
		token, s, err := utils.IsValidJwtAccessToken(app, tokenString)
		if err != nil {
			context.Status(http.StatusBadRequest)
			context.Abort()
			return
		}
		if !token {
			context.Status(http.StatusBadRequest)
			context.Abort()
			return
		}
		context.Set("userId", s)
		context.Next()
	}

}
