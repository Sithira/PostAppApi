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

func (c commentRepository) GetComment(ctx context.Context, postId uuid.UUID, commentId uuid.UUID) (*entities.Comment, error) {
	statement, err := c.db.PrepareContext(ctx, comments.GetCommentForPost)
	if err != nil {
		return nil, err
	}

	cmt := &entities.Comment{}

	if err := statement.QueryRowContext(ctx, postId.String(), commentId.String()).Scan(&cmt.ID, &cmt.CommentBody, &cmt.CreatedAt, &cmt.UpdatedAt, &cmt.DeletedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(err, "no results")
		}
	}

	return cmt, nil
}

func (c commentRepository) GetCommentsForPostId(ctx context.Context, postId uuid.UUID) ([]*entities.Comment, error) {
	statement, err := c.db.PrepareContext(ctx, comments.GetAllCommentsForPost)

	if err != nil {
		return nil, err
	}

	var cmts []*entities.Comment

	results, err := statement.QueryContext(ctx, postId)

	for results.Next() {
		var comment entities.Comment
		if err := results.Scan(&comment.ID, &comment.PostId, &comment.ParentCommentId, &comment.CommentBody, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
			return nil, errors.Wrap(err, "comments_repo.GetCommentsForPostId")
		}
		cmts = append(cmts, &comment)
	}

	return cmts, nil
}

func (c commentRepository) AddCommentForPost(ctx context.Context, commentBody string, postId uuid.UUID, userId uuid.UUID) (*entities.Comment, error) {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	statement, err := tx.PrepareContext(ctx, comments.InsertCommentForPost)

	if err != nil {
		_ = tx.Rollback()
		return nil, errors.Wrap(err, "comments_repo.AddCommentForPost")
	}

	defer statement.Close()

	newComment := &entities.Comment{}

	if err := statement.QueryRowContext(ctx, uuid.New(), postId, userId, commentBody).Scan(&newComment.ID); err != nil {
		_ = tx.Rollback()
		return nil, errors.Wrap(err, "comments_repo.AddCommentForPost")
	}

	_ = tx.Commit()
	return newComment, nil
}
func (c commentRepository) UpdateCommentForPost(ctx context.Context, comment entities.Comment) error {
	tx, err := c.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	results, err := tx.ExecContext(ctx, comments.UpdateCommentForPost, comment.CommentBody, comment.ID.String())

	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return errors.Wrap(err, "comments_repo.UpdateCommentForPost.update.Exec")
	}

	affected, err := results.RowsAffected()

	if err != nil {
		_ = tx.Rollback()
		return errors.Wrap(err, "comments_repo.UpdateCommentForPost.RowsAffected")
	}
	if affected > 0 {
		_ = tx.Commit()
		return nil
	}

	err = tx.Rollback()
	if err != nil {
		return err
	}
	return errors.New("update not performed")
}

func (c commentRepository) DeleteComment(ctx context.Context, postId uuid.UUID, commentId uuid.UUID) error {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	results, err := tx.ExecContext(ctx, comments.DeleteCommentForPost, postId.String(), commentId.String())
	affected, err := results.RowsAffected()

	if affected > 0 {
		err := tx.Commit()
		if err != nil {
			return err
		}
		return nil
	}

	err = tx.Rollback()

	if err != nil {
		return err
	}

	return errors.New("delete not performed")
}
