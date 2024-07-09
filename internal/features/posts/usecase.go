package posts

import (
	"RestApiBackend/internal/features/posts/dto"
	"context"
)

type UseCase interface {
	FetchPosts(ctx context.Context) (*dto.PostsListResponse, error)

	CreatePost(ctx context.Context, comment *dto.CreatePostRequest) (*dto.CreatePostResponse, error)

	UpdatePost(ctx context.Context, postId string, comment *dto.UpdatePostRequest) (*dto.CreatePostResponse, error)
}
