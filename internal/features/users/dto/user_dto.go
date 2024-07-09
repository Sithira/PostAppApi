package dto

import (
	"github.com/google/uuid"
	"time"
)

type UserCreateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserCreateResponse struct {
	UserID       int64  `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiredAt    int64  `json:"expired_at"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponseList struct {
	TotalCount int             `json:"total_count"`
	TotalPages int             `json:"total_pages"`
	Page       int             `json:"page"`
	Size       int             `json:"size"`
	HasMore    bool            `json:"has_more"`
	Users      []*UserResponse `json:"users"`
}
