package http

import (
	"RestApiBackend/internal/features/auth"
	"RestApiBackend/internal/features/auth/dto"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

type authHandler struct {
	uc auth.AuthenticationUseCase
}

func NewAuthenticationHandler(useCase auth.AuthenticationUseCase) auth.AuthenticationHandler {
	return &authHandler{uc: useCase}
}

func (auth authHandler) Register() gin.HandlerFunc {
	return func(context *gin.Context) {
		var registerRequest *dto.RegisterRequest
		err := context.BindJSON(&registerRequest)
		fmt.Printf("Register request body = %v %T", registerRequest, registerRequest)

		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		register, err := auth.uc.Register(context, registerRequest)

		if err != nil {
			context.JSON(http.StatusBadRequest, errors.Wrap(err, "unable to register user"))
			return
		}
		context.JSON(http.StatusOK, register)
		return
	}
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
