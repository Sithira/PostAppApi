package comments

import (
	"RestApiBackend/internal/features/comments/dto"
	"context"
	"github.com/google/uuid"
)

type UseCase interface {
	FetchCommentsForPost(ctx context.Context, userId uuid.UUID, postId uuid.UUID) (*dto.CommentListResponse, error)

	AddCommentForPost(ctx context.Context, userId uuid.UUID, postId uuid.UUID, request dto.AddCommentRequest) (*dto.CommentResponse, error)

	FetchComment(ctx context.Context, postId uuid.UUID, commentId uuid.UUID) (*dto.CommentResponse, error)

	UpdateComment(ctx context.Context, postId uuid.UUID, commentId uuid.UUID) error

	DeleteComment(ctx context.Context, postId uuid.UUID)
}
