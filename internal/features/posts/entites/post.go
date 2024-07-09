package entites

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	ID        uuid.UUID `db:"id"`
	UserId    uuid.UUID `db:"user_id"`
	Title     string    `db:"title"`
	Body      string    `db:"body"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"updated_at"`
}

type UserPosts struct {
	PostId uuid.UUID `db:"post_id"`
	UserId uuid.UUID `db:"user_id"`
}
