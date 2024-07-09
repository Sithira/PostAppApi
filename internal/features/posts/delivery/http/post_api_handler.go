package http

import (
	"RestApiBackend/internal/features/posts"
	"RestApiBackend/internal/features/posts/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type postUseCase struct {
	us posts.UseCase
}

func NewPostHandler(uc posts.UseCase) posts.Handler {
	return &postUseCase{
		us: uc,
	}
}

func (p postUseCase) GetPostsForUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		userId := context.Query("userId")
		context.Set("userId", userId)
		fetchPosts, err := p.us.FetchPosts(context)
		if err != nil {
			context.JSON(http.StatusBadRequest, nil)
		}
		context.JSON(http.StatusOK, fetchPosts)
		return
	}
}

func (p postUseCase) CreatePostForUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		userId := context.Query("userId")
		context.Set("userId", userId)
		var body *dto.CreatePostRequest
		err := context.BindJSON(&body)
		if err != nil {
			context.JSON(400, "Unable to deser body")
			return
		}
		post, err := p.us.CreatePost(context, body)
		if err != nil {
			context.JSON(400, err)
			return
		}
		context.JSON(http.StatusOK, post)
		return
	}
}

func (p postUseCase) UpdatePostForUser() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (p postUseCase) DeletePostForUser() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (p postUseCase) AttachPostToUser() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (p postUseCase) DeAttachPostFromUser() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}
