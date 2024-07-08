package http

import (
	"RestApiBackend/internal/features/users"
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
		email := context.Query("email")
		if len(email) <= 0 {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
			return
		}
		username, err := uc.userUseCase.GetUserByEmail(context, email)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "User retrieved successfully",
			"data":    username,
		})
	}
}
