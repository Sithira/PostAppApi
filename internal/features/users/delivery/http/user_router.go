package http

import (
	"RestApiBackend/infrastructure"
	"RestApiBackend/internal/features/users"
	"github.com/gin-gonic/gin"
)

func UserRoutes(app *infrastructure.Application, h users.Handlers, group *gin.RouterGroup) {
	group.GET("", h.GetUserDetails())
}
