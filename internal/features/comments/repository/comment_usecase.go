package repository

import (
	"RestApiBackend/internal/features/comments"
	"RestApiBackend/internal/features/comments/dto"
	"RestApiBackend/internal/features/posts"
	"context"
	"github.com/google/uuid"
)

type commentsUseCase struct {
	commentRepo comments.CommentRepository
	postRepo    posts.PostRepository
}

func NewCommentsUseCase(cRepo comments.CommentRepository, pRepo posts.PostRepository) comments.UseCase {
	return &commentsUseCase{
		commentRepo: cRepo,
		postRepo:    pRepo,
	}
}

func (c commentsUseCase) FetchComment(ctx context.Context, postId uuid.UUID, commentId uuid.UUID) (*dto.CommentResponse, error) {
	c.commentRepo.GetComment(ctx, postId, commentId)
}

func (c commentsUseCase) FetchCommentsForPost(ctx context.Context, postId uuid.UUID) (*dto.CommentListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c commentsUseCase) UpdateComment(ctx context.Context, userId uuid.UUID, postId uuid.UUID, commentId uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (c commentsUseCase) DeleteComment(ctx context.Context, userId uuid.UUID, postId uuid.UUID) {
	//TODO implement me
	panic("implement me")
}

func (c commentsUseCase) AddCommentForPost(ctx context.Context, userId uuid.UUID, postId uuid.UUID, request dto.AddCommentRequest) (*dto.CommentResponse, error) {
	//TODO implement me
	panic("implement me")
}
