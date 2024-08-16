package http

import (
	"RestApiBackend/internal/features/comments"
	"RestApiBackend/internal/features/comments/dto"
	httperror "RestApiBackend/pkg/http"
	"RestApiBackend/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

type commentApiHandler struct {
	uc comments.UseCase
}

func NewCommentApiHandler(commentsUseCase comments.UseCase) comments.Handler {
	return &commentApiHandler{
		uc: commentsUseCase,
	}
}

func (c commentApiHandler) GetSingleComment() gin.HandlerFunc {
	return func(context *gin.Context) {
		utils.GetUserDetailsFromContext(context)
		postUuid, commentUuid := extractParams(context)
		comment, err := c.uc.FetchComment(context, *postUuid, *commentUuid)
		if err != nil {
			return
		}
		context.JSON(http.StatusOK, comment)
		return
	}
}

func (c commentApiHandler) GetComments() gin.HandlerFunc {
	return func(context *gin.Context) {
		utils.GetUserDetailsFromContext(context)
		postUuid, _ := extractParams(context)
		postList, err := c.uc.FetchCommentsForPost(context, *postUuid)
		if err != nil {
			return
		}
		context.JSON(http.StatusOK, postList)
		return
	}
}

func (c commentApiHandler) AddComment() gin.HandlerFunc {
	return func(context *gin.Context) {
		_, user := utils.GetUserDetailsFromContext(context)
		postUuid, _ := extractParams(context)

		var body *dto.AddCommentRequest
		err := context.ShouldBindJSON(&body)

		if err != nil {
			return
		}

		post, err := c.uc.AddCommentForPost(context, user.ID, *postUuid, *body)
		if err != nil {
			return
		}

		context.JSON(http.StatusOK, post)
		return
	}
}

func (c commentApiHandler) UpdateComment() gin.HandlerFunc {
	return func(context *gin.Context) {
		_, user := utils.GetUserDetailsFromContext(context)
	}
}

func (c commentApiHandler) DeleteComment() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (c commentApiHandler) AddCommentForComment() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func extractParams(context *gin.Context) (*uuid.UUID, *uuid.UUID) {
	postId := strings.TrimSpace(context.Param("postId"))
	if postId == "" || len(postId) == 0 {
		context.AbortWithStatus(http.StatusBadRequest)
		return nil, nil
	}

	uuidPostId, err := uuid.Parse(postId)
	if err != nil {
		context.JSON(httperror.ErrorResponse(err))
		return nil, nil
	}

	commentId := strings.TrimSpace(context.Param("commentId"))

	if commentId != "" || len(commentId) != 0 {
		uuidCommentId, err := uuid.Parse(commentId)
		if err != nil {
			context.JSON(httperror.ErrorResponse(err))
			return nil, nil
		}
		return &uuidPostId, &uuidCommentId
	}

	return &uuidPostId, nil
}
