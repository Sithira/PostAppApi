package repository

import (
	"RestApiBackend/internal/features/posts"
	"RestApiBackend/internal/features/posts/entites"
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
	statement, err := p.db.PrepareContext(ctx, posts.InsertPostByUserId)

	if err != nil {
		return nil, errors.Wrap(err, "post_repo.CreatePostForUser.prepare")
	}

	defer statement.Close()

	pst := &entites.Post{}

	if err := statement.QueryRowContext(ctx, uuid.New(), userId, request.Title, request.Body).Scan(
		&pst.ID,
	); err != nil {
		return nil, errors.Wrap(err, "post_repo.CreatePostForUser.insert.Scan")
	}

	return pst, nil
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
