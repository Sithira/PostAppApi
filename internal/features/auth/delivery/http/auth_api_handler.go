package http

import (
	"RestApiBackend/internal/features/auth"
	"RestApiBackend/internal/features/auth/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authHandler struct {
	uc auth.AuthenticationUseCase
}

func NewAuthenticationHandler(useCase auth.AuthenticationUseCase) auth.AuthenticationHandler {
	return &authHandler{uc: useCase}
}

func (auth authHandler) Login() gin.HandlerFunc {
	return func(context *gin.Context) {
		var loginRequest dto.LoginRequest
		err := context.BindJSON(&loginRequest)
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		login, userId, err := auth.uc.Login(context.Request.Context(), &loginRequest)
		if err != nil {
			context.AbortWithStatus(http.StatusForbidden)
			return
		}
		context.Set("userId", userId)
		context.JSON(http.StatusOK, login)
		return
	}
}
