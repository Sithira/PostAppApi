package entities

import (
	"github.com/google/uuid"
	"time"
)

type Comment struct {
	ID              uuid.UUID `db:"id"`
	PostId          uuid.UUID `db:"post_id"`
	ParentCommentId uuid.UUID `db:"post_id"`
	CommentBody     string    `db:"comment_body"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
	DeletedAt       time.Time `db:"updated_at"`
}

func NewComment() *Comment {
	return &Comment{ID: uuid.New()}
}
