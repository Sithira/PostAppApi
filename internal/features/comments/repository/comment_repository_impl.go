package repository

import (
	"RestApiBackend/internal/features/comments"
	"RestApiBackend/internal/features/comments/entities"
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type commentRepository struct {
	db *sql.DB
}

func NewCommentRepository(conn *sql.DB) comments.CommentRepository {
	return &commentRepository{
		db: conn,
	}
}

func (c commentRepository) GetComment(ctx context.Context, userId uuid.UUID, postId uuid.UUID, commentId uuid.UUID) (*entities.Comment, error) {
	statement, err := c.db.PrepareContext(ctx, comments.GetCommentForPost)
	if err != nil {
		return nil, err
	}

	cmt := &entities.Comment{}

	if err := statement.QueryRowContext(ctx, postId.String()).Scan(&cmt.ID, &cmt.CommentBody, &cmt.CreatedAt, &cmt.UpdatedAt, &cmt.DeletedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(err, "no results")
		}
	}

	return cmt, nil
}

func (c commentRepository) GetCommentsForPostId(ctx context.Context, postId uuid.UUID) ([]*entities.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (c commentRepository) AddCommentForPost(ctx context.Context, commentBody string, postId uuid.UUID) (*entities.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (c commentRepository) UpdateCommentForPost(ctx context.Context, commentBody string, postId uuid.UUID, commentId uuid.UUID) (*entities.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (c commentRepository) DeleteComment(ctx context.Context, postId uuid.UUID, commentId uuid.UUID) {
	//TODO implement me
	panic("implement me")
}
