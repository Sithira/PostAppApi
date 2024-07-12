package posts

import (
	"RestApiBackend/internal/features/posts/dto"
	"github.com/gin-gonic/gin"
)

type UseCase interface {
	FetchPosts(ctx *gin.Context) (*dto.PostsListResponse, error)

	CreatePost(ctx *gin.Context, comment *dto.CreatePostRequest) (*dto.CreatePostResponse, error)

	UpdatePost(ctx *gin.Context, postId string, comment *dto.UpdatePostRequest) (*dto.CreatePostResponse, error)
}
