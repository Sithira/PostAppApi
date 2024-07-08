package server

import (
	"RestApiBackend/infrastructure"
	"RestApiBackend/internal/features/users/delivery/http"
	"RestApiBackend/internal/features/users/repository"
	"RestApiBackend/internal/features/users/usecase"
	"github.com/gin-gonic/gin"
)

func (s *Server) MapHandlers() {
	// invoke repositories
	userRepository := repository.NewUserRepository(s.db)

	// use cases
	userUc := usecase.NewUserUserCase(userRepository)

	// handlers
	userHandler := http.NewUserHandler(userUc)

	// base url of the application
	baseRouter := s.gin.Group("")

	PingRoute(s.app, baseRouter)

	// register routes
	http.UserRoutes(s.app, userHandler, baseRouter.Group("/api/v1/users"))
}

func PingRoute(app *infrastructure.Application, router *gin.RouterGroup) {
	router.GET("ping", func(context *gin.Context) {
		context.String(200, "pong")
	})
}
