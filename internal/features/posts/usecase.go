package posts

import (
	"RestApiBackend/internal/features/posts/dto"
	"context"
	"github.com/google/uuid"
)

type UseCase interface {
	FetchPosts(ctx context.Context, userId uuid.UUID) (*dto.PostsListResponse, error)

	FetchPost(ctx context.Context, userId uuid.UUID, postId uuid.UUID) (*dto.PostResponse, error)

	CreatePost(ctx context.Context, userId uuid.UUID, postRequest dto.CreatePostRequest) (*dto.CreatePostResponse, error)

	UpdatePost(ctx context.Context, userId uuid.UUID, postId uuid.UUID, postRequest dto.UpdatePostRequest) error

	DeletePost(ctx context.Context, userId uuid.UUID, postId uuid.UUID) error
}
