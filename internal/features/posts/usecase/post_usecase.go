package usecase

import (
	"RestApiBackend/internal/features/posts"
	"RestApiBackend/internal/features/posts/dto"
	"RestApiBackend/internal/features/posts/entites"
	"RestApiBackend/pkg/http"
	"context"
	"github.com/google/uuid"
)

type postUseCase struct {
	postRepository posts.PostRepository
}

func NewPostUseCase(repository posts.PostRepository) posts.UseCase {
	return &postUseCase{
		postRepository: repository,
	}
}

func (p postUseCase) FetchPosts(ctx context.Context) (*dto.PostsListResponse, error) {
	var userIdFromCtx = ctx.Value("userId").(string)
	userId, err := uuid.Parse(userIdFromCtx)
	if err != nil {

	}
	fetchedPosts, err := p.postRepository.FetchPostsOfUser(ctx, userId)
	if err != nil {
		return nil, http.NewBadRequest(err)
	}
	return toPostResponseList(fetchedPosts), nil
}

func (p postUseCase) CreatePost(ctx context.Context, post *dto.CreatePostRequest) (*dto.CreatePostResponse, error) {
	var userIdFromCtx = ctx.Value("userId").(string)
	userId, err := uuid.Parse(userIdFromCtx)

	if err != nil {
		return nil, http.NewBadRequest(err)
	}

	duplicateExists, err := p.postRepository.FindDuplicatedByPostTitle(ctx, post.Title, userId)
	if err != nil {
		return nil, http.NewInternalServerError(err)
	}

	if *duplicateExists {
		return nil, http.NewBadRequest("Duplicate post exists")
	}

	createdPost, err := p.postRepository.CreatePostForUser(ctx, userId, post)

	if err != nil {
		return nil, http.NewBadRequest(err)
	}

	return convertToCreatedPostResponse(createdPost), nil
}

func convertToCreatedPostResponse(post *entites.Post) *dto.CreatePostResponse {
	return &dto.CreatePostResponse{ID: post.ID}
}

func (p postUseCase) UpdatePost(ctx context.Context, postId string, comment *dto.UpdatePostRequest) (*dto.CreatePostResponse, error) {
	return nil, nil
}

func toPostResponseList(p []*entites.Post) *dto.PostsListResponse {
	postList := make([]*dto.PostResponse, len(p))
	for i, post := range p {
		postList[i] = convertUserDTOToUser(post)
	}
	return &dto.PostsListResponse{
		TotalCount: len(postList),
		Page:       0,
		Size:       0,
		HasMore:    false,
		Posts:      postList,
	}
}

func convertUserDTOToUser(post *entites.Post) *dto.PostResponse {
	return &dto.PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		BodyText:  post.Body,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}
