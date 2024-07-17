package posts

import (
	"RestApiBackend/internal/features/posts/entites"
	"context"
	"github.com/google/uuid"
)

const (
	SelectPostById               = `SELECT p.id, p.user_id, p.title, p.body, p.created_at, p.updated_at FROM posts p WHERE p.id = $1 AND p.user_id =$2 AND p.deleted_at IS NULL`
	SelectPostsByUserId          = `SELECT p.id, p.title, p.body, p.created_at, p.updated_at FROM posts p WHERE p.user_id = $1 AND p.deleted_at IS NULL`
	InsertPostByUserId           = `INSERT INTO posts (id, user_id, title, body, created_at, updated_at) VALUES ($1, $2, $3, $4, now(), now()) RETURNING id`
	UpdatePostOfUser             = `UPDATE posts SET title = NULLIF($3, ''), body  = NULLIF($4, ''), updated_at = now() WHERE id = $1 AND user_id = $2 AND deleted_at is NULL;`
	DeletePostOfUser             = `UPDATE posts SET deleted_at = now() WHERE id = $1 AND user_id = $2 AND deleted_at is NULL;`
	FindDuplicateRecordsByUserId = `SELECT COUNT(p.id) > 0  FROM posts p WHERE p.title = $1 AND p.user_id = $2 AND p.deleted_at IS NULL`
)

type PostRepository interface {
	FetchPost(ctx context.Context, userId uuid.UUID, id uuid.UUID) (*entites.Post, error)

	FetchPostsOfUser(ctx context.Context, userId uuid.UUID) ([]*entites.Post, error)

	FindDuplicatedByPostTitle(ctx context.Context, postTitle string, userId uuid.UUID) (*bool, error)

	CreatePostForUser(ctx context.Context, userId uuid.UUID, request entites.Post) (*entites.Post, error)

	UpdatePostOfUser(ctx context.Context, request entites.Post) error

	DeletePostOfUser(ctx context.Context, request entites.Post) error
}
