package comments

import (
	"RestApiBackend/internal/features/comments/entities"
	"context"
	"github.com/google/uuid"
)

const (
	GetCommentForPost     = "SELECT c.* FROM posts p JOIN comments c ON p.id = c.post_id WHERE p.id = $1 AND c.id = $2 AND (p.deleted_at IS NULL AND c.deleted_at IS NULL)"
	GetAllCommentsForPost = "SELECT * FROM posts p JOIN public.comments c on p.id = c.post_id WHERE p.id = $1 AND p.deleted_at AND c.deleted_at IS NULL"
	InsertCommentForPost  = "INSERT INTO comments (id, post_id, user_id, comment_body, created_at, updated_at) VALUES ($1, $2, $3, $4, now(), now()) RETURNING id"
	UpdateCommentForPost  = "UPDATE comments c SET c.comment_body = $1 WHERE c.id = $2 AND c.deleted_at IS NULL"
	DeleteCommentForPost  = ""
)

type CommentRepository interface {
	GetComment(ctx context.Context, postId uuid.UUID, commentId uuid.UUID) (*entities.Comment, error)

	GetCommentsForPostId(ctx context.Context, postId uuid.UUID) ([]*entities.Comment, error)

	AddCommentForPost(ctx context.Context, commentBody string, postId uuid.UUID, userId uuid.UUID) (*entities.Comment, error)

	UpdateCommentForPost(ctx context.Context, comment entities.Comment) error

	DeleteComment(ctx context.Context, postId uuid.UUID, commentId uuid.UUID) error
}
