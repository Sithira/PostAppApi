package comments

import (
	"RestApiBackend/internal/features/comments/entities"
	"context"
	"github.com/google/uuid"
)

const (
	GetCommentForPost     = ""
	GetAllCommentsForPost = ""
	InsertCommentForPost  = ""
	UpdateCommentForPost  = ""
	DeleteCommentForPost  = ""
)

type CommentRepository interface {
	GetComment(ctx context.Context, userId uuid.UUID, postId uuid.UUID, commentId uuid.UUID) (*entities.Comment, error)

	GetCommentsForPostId(ctx context.Context, postId uuid.UUID) ([]*entities.Comment, error)

	AddCommentForPost(ctx context.Context, commentBody string, postId uuid.UUID) (*entities.Comment, error)

	UpdateCommentForPost(ctx context.Context, commentBody string, postId uuid.UUID, commentId uuid.UUID) (*entities.Comment, error)

	DeleteComment(ctx context.Context, postId uuid.UUID, commentId uuid.UUID)
}
