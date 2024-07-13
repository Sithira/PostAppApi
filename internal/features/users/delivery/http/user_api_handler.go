package http

import (
	"RestApiBackend/internal/features/users"
	"RestApiBackend/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHandler struct {
	userUseCase users.UseCase
}

func NewUserHandler(userUc users.UseCase) users.Handlers {
	return &userHandler{
		userUseCase: userUc,
	}
}

func (uc *userHandler) GetUserDetails() gin.HandlerFunc {
	return func(context *gin.Context) {
		_, userDetails := utils.GetUserDetailsFromContext(context)
		username, err := uc.userUseCase.GetUserByEmail(context.Request.Context(), userDetails.Email)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		context.JSON(http.StatusOK, username)
	}
}
