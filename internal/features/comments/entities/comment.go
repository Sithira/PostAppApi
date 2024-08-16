package entities

import (
	"github.com/google/uuid"
	"time"
)

type Comment struct {
	ID              uuid.UUID `db:"id"`
	UserId          uuid.UUID `db:"user_id"`
	PostId          uuid.UUID `db:"post_id"`
	ParentCommentId uuid.UUID `db:"parent_comment_id"`
	CommentBody     string    `db:"comment_body"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
	DeletedAt       time.Time `db:"updated_at"`
}

func NewComment(postId uuid.UUID, userId uuid.UUID) *Comment {
	return &Comment{ID: uuid.New(), PostId: postId, UserId: userId}
}
