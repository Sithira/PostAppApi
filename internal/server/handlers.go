package server

import (
	"RestApiBackend/infrastructure"
	http2 "RestApiBackend/internal/features/posts/delivery/http"
	repository2 "RestApiBackend/internal/features/posts/repository"
	usecase2 "RestApiBackend/internal/features/posts/usecase"
	"RestApiBackend/internal/features/users/delivery/http"
	"RestApiBackend/internal/features/users/repository"
	"RestApiBackend/internal/features/users/usecase"
	"github.com/gin-gonic/gin"
)

func (s *Server) MapHandlers() {
	// invoke repositories
	postRepository := repository2.NewPostsRepository(s.db)
	userRepository := repository.NewUserRepository(s.db)

	// use cases
	userUc := usecase.NewUserUserCase(userRepository)
	postUc := usecase2.NewPostUseCase(postRepository)

	// handlers
	userHandler := http.NewUserHandler(userUc)
	postHandler := http2.NewPostHandler(postUc)

	// base url of the application
	baseRouter := s.gin.Group("")

	PingRoute(s.app, baseRouter)

	// register routes
	http.UserRoutes(s.app, userHandler, baseRouter.Group("/api/v1/users"))
	http2.NewPostRouter(postHandler, baseRouter.Group("/api/v1/posts"))
}

func PingRoute(app *infrastructure.Application, router *gin.RouterGroup) {
	router.GET("ping", func(context *gin.Context) {
		context.String(200, "pong")
	})
}
