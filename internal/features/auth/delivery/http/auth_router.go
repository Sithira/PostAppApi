package http

import (
	"RestApiBackend/internal/features/auth"
	"github.com/gin-gonic/gin"
)

func NewAuthRouter(authHandler auth.AuthenticationHandler, group *gin.RouterGroup) {
	group.POST("/login", authHandler.Login())
	group.POST("/register", authHandler.Register())
}
