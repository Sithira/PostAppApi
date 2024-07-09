package posts

import (
	"RestApiBackend/internal/features/posts/dto"
	"RestApiBackend/internal/features/posts/entites"
	"context"
	"github.com/google/uuid"
)

const (
	SelectPostsByUserId          = `SELECT p.id, p.title, p.body, p.created_at, p.updated_at FROM posts p WHERE p.user_id = $1 AND p.deleted_at IS NULL`
	InsertPostByUserId           = `INSERT INTO posts (id, user_id, title, body, created_at, updated_at) VALUES ($1, $2, $3, $4, now(), now()) RETURNING id`
	FindDuplicateRecordsByUserId = `SELECT COUNT(p.id) > 0  FROM posts p WHERE p.title = $1 AND p.user_id = $2 AND p.deleted_at IS NULL`
)

type PostRepository interface {
	CreatePostForUser(ctx context.Context, userId uuid.UUID, request *dto.CreatePostRequest) (*entites.Post, error)

	FindDuplicatedByPostTitle(ctx context.Context, postTitle string, userId uuid.UUID) (*bool, error)

	FetchPostsOfUser(ctx context.Context, userId uuid.UUID) ([]*entites.Post, error)
}
