package dto

import (
	"github.com/google/uuid"
	"time"
)

type CreatePostRequest struct {
	Title    string `json:"title"`
	BodyText string `json:"body_text"`
	Tags     string `json:"tags"`
}

type UpdatePostRequest struct {
	Title    string `json:"title"`
	BodyText string `json:"body_text"`
}

type CreatePostResponse struct {
	ID uuid.UUID `json:"id"`
}

type PostResponse struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	BodyText  string    `json:"body_text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostsListResponse struct {
	TotalCount int             `json:"total_count"`
	TotalPages int             `json:"total_pages"`
	Page       int             `json:"page"`
	Size       int             `json:"size"`
	HasMore    bool            `json:"has_more"`
	Data       []*PostResponse `json:"data"`
}
