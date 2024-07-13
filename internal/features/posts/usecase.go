package posts

import (
	"RestApiBackend/internal/features/posts/dto"
	"context"
	"github.com/google/uuid"
)

type UseCase interface {
	FetchPosts(ctx context.Context, userId uuid.UUID) (*dto.PostsListResponse, error)

	CreatePost(ctx context.Context, userId uuid.UUID, comment *dto.CreatePostRequest) (*dto.CreatePostResponse, error)

	UpdatePost(ctx context.Context, userId uuid.UUID, postId string, comment *dto.UpdatePostRequest) (*dto.CreatePostResponse, error)
}
