package usecase

import (
	"RestApiBackend/internal/features/posts"
	"RestApiBackend/internal/features/posts/dto"
	"RestApiBackend/internal/features/posts/entites"
	"RestApiBackend/pkg/http"
	"RestApiBackend/pkg/utils"
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
)

type postUseCase struct {
	postRepository posts.PostRepository
}

func NewPostUseCase(repository posts.PostRepository) posts.UseCase {
	return &postUseCase{
		postRepository: repository,
	}
}

func (p postUseCase) FetchPosts(ctx context.Context, userId uuid.UUID) (*dto.PostsListResponse, error) {
	fetchedPosts, err := p.postRepository.FetchPostsOfUser(ctx, userId)
	if err != nil {
		return nil, http_error.InternalServerError
	}
	return toPostResponseList(fetchedPosts), nil
}

func (p postUseCase) CreatePost(ctx context.Context, userId uuid.UUID, post dto.CreatePostRequest) (*dto.CreatePostResponse, error) {
	validate := validator.New()

	if err := validate.Struct(post); err != nil {
		var errs validator.ValidationErrors
		errors.As(err, &errs)
		errorList := make(map[string]string)
		for _, fieldErr := range errs {
			errorList[fieldErr.Field()] = fieldErr.Tag()
		}
		return nil, http_error.NewRestError(http.StatusBadRequest, "ERR_000", errorList)
	}

	duplicateExists, err := p.postRepository.FindDuplicatedByPostTitle(ctx, utils.ToString(post.Title), userId)

	if err != nil {
		return nil, http_error.NewInternalServerError(err)
	}

	if *duplicateExists {
		return nil, http_error.NewRestError(http.StatusBadRequest, "ERR_001", "Duplicate")
	}

	postEntity := entites.NewPost()
	postEntity.Title = utils.ToString(post.Title)
	postEntity.Body = utils.ToString(post.BodyText)

	createdPost, err := p.postRepository.CreatePostForUser(ctx, userId, *postEntity)

	if err != nil {
		return nil, http_error.NewRestError(http.StatusInternalServerError, "", "")
	}

	return convertToCreatedPostResponse(createdPost), nil
}

func (p postUseCase) UpdatePost(ctx context.Context, userId uuid.UUID, postId string, comment *dto.UpdatePostRequest) (*dto.CreatePostResponse, error) {
	//TODO implement me
	panic("implement me")
}

func convertToCreatedPostResponse(post *entites.Post) *dto.CreatePostResponse {
	return &dto.CreatePostResponse{ID: post.ID}
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
		Data:       postList,
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
