package server

import (
	"RestApiBackend/infrastructure"
	http3 "RestApiBackend/internal/features/auth/delivery/http"
	usecase3 "RestApiBackend/internal/features/auth/usecase"
	http4 "RestApiBackend/internal/features/comments/delivery/http"
	repository3 "RestApiBackend/internal/features/comments/repository"
	http2 "RestApiBackend/internal/features/posts/delivery/http"
	repository2 "RestApiBackend/internal/features/posts/repository"
	usecase2 "RestApiBackend/internal/features/posts/usecase"
	"RestApiBackend/internal/features/users/delivery/http"
	"RestApiBackend/internal/features/users/repository"
	"RestApiBackend/internal/features/users/usecase"
	"RestApiBackend/internal/server/middlewares"
	"RestApiBackend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (s *Server) MapHandlers() {
	// invoke repositories
	postRepository := repository2.NewPostsRepository(s.db)
	userRepository := repository.NewUserRepository(s.db)
	commentsRepository := repository3.NewCommentRepository(s.db)

	// use cases
	userUc := usecase.NewUserUserCase(userRepository)
	postUc := usecase2.NewPostUseCase(postRepository)
	commentUs := repository3.NewCommentsUseCase(commentsRepository, postRepository)
	authUs := usecase3.NewAuthenticationUseCase(*s.app, userRepository)

	// handlers
	userHandler := http.NewUserHandler(userUc)
	postHandler := http2.NewPostHandler(postUc)
	commentsHandler := http4.NewCommentApiHandler(commentUs)
	authHandler := http3.NewAuthenticationHandler(authUs)

	// base url of the application
	baseRouter := s.gin.Group("")

	PingRoute(s.app, baseRouter)

	// register routes
	bearerToken := middlewares.NewAuthBearerToken(userRepository, s.app, utils.IsValidJwtAccessToken)

	userRouter := baseRouter.Group("/api/v1/users")
	userRouter.Use(bearerToken)
	http.UserRoutes(s.app, userHandler, userRouter)

	postsGroup := baseRouter.Group("/api/v1/posts")
	postsGroup.Use(bearerToken)

	commentsGroup := baseRouter.Group("/api/v1/posts/:postId/comments/")
	commentsGroup.Use(bearerToken)

	http2.NewPostRouter(postHandler, postsGroup)
	http4.NewCommentRouter(commentsHandler, commentsGroup)
	http3.NewAuthRouter(authHandler, baseRouter.Group("/api/auth"))
}

func PingRoute(app *infrastructure.Application, router *gin.RouterGroup) {
	router.GET("ping", func(context *gin.Context) {
		context.String(200, "pong")
	})
}
