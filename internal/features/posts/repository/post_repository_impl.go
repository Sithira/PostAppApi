package repository

import (
	"RestApiBackend/internal/features/posts"
	"RestApiBackend/internal/features/posts/entites"
	httperror "RestApiBackend/pkg/http"
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type postRepository struct {
	db *sql.DB
}

func NewPostsRepository(db *sql.DB) posts.PostRepository {
	return &postRepository{
		db: db,
	}
}

func (p postRepository) FetchPost(ctx context.Context, userId uuid.UUID, postId uuid.UUID) (*entites.Post, error) {
	statement, err := p.db.PrepareContext(ctx, posts.SelectPostById)

	if err != nil {
		return nil, errors.Wrap(err, "query prepare failure")
	}

	defer statement.Close()

	post := &entites.Post{}

	if err := statement.QueryRowContext(ctx, postId.String(), userId.String()).Scan(&post.ID, &post.UserId, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(err, "no results")
		}
		return nil, httperror.NewInternalServerError(nil)
	}
	return post, nil
}

func (p postRepository) FetchPostsOfUser(ctx context.Context, userId uuid.UUID) ([]*entites.Post, error) {
	statement, err := p.db.PrepareContext(ctx, posts.SelectPostsByUserId)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	results, err := statement.QueryContext(ctx, userId.String())

	if err != nil {
		return nil, err
	}

	var fetchedPosts []*entites.Post

	for results.Next() {
		var fetchedPost entites.Post
		if err := results.Scan(&fetchedPost.ID, &fetchedPost.Title, &fetchedPost.Body, &fetchedPost.CreatedAt, &fetchedPost.UpdatedAt); err != nil {
			return nil, errors.Wrap(err, "post_repo.FetchPostsOfUser")
		}
		fetchedPosts = append(fetchedPosts, &fetchedPost)
	}

	if err = results.Err(); err != nil {
		return nil, err
	}

	return fetchedPosts, nil
}

func (p postRepository) CreatePostForUser(ctx context.Context, userId uuid.UUID, request entites.Post) (*entites.Post, error) {
	tx, err := p.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	statement, err := tx.PrepareContext(ctx, posts.InsertPostByUserId)

	if err != nil {
		_ = tx.Rollback()
		return nil, errors.Wrap(err, "post_repo.CreatePostForUser.prepare")
	}

	defer statement.Close()

	pst := &entites.Post{}

	if err := statement.QueryRowContext(ctx, uuid.New(), userId, request.Title, request.Body).Scan(
		&pst.ID,
	); err != nil {
		_ = tx.Rollback()
		return nil, errors.Wrap(err, "post_repo.CreatePostForUser.insert.Scan")
	}

	_ = tx.Commit()
	return pst, nil
}

func (p postRepository) UpdatePostOfUser(ctx context.Context, request entites.Post) error {
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	result, err := tx.ExecContext(ctx, posts.UpdatePostOfUser, request.ID.String(), request.UserId.String(), request.Title, request.Body)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return errors.Wrap(err, "post_repo.UpdatePostOfUser.update.Exec")
	}
	affected, err := result.RowsAffected()
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return errors.Wrap(err, "post_repo.UpdatePostOfUser.update.Exec.RowsAffected")
	}
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
	return errors.New("update not performed")
}

func (p postRepository) DeletePostOfUser(ctx context.Context, request entites.Post) error {
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	result, err := tx.ExecContext(ctx, posts.DeletePostOfUser, request.ID.String(), request.UserId.String())
	affected, err := result.RowsAffected()

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

func (p postRepository) FindDuplicatedByPostTitle(ctx context.Context, postTitle string, userId uuid.UUID) (*bool, error) {
	statement, err := p.db.PrepareContext(ctx, posts.FindDuplicateRecordsByUserId)

	if err != nil {
		return nil, errors.Wrap(err, "post_repo.FindDuplicatedByPostTitle.prepare")
	}

	defer statement.Close()

	var duplicateExists bool

	if err := statement.QueryRowContext(ctx, postTitle, userId).Scan(&duplicateExists); err != nil {
		return nil, errors.Wrap(err, "post_repo.FindDuplicatedByPostTitle.scan")
	}

	return &duplicateExists, err
}
