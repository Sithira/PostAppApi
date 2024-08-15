package comments

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetSingleComment() gin.HandlerFunc

	GetComments() gin.HandlerFunc

	AddComment() gin.HandlerFunc

	UpdateComment() gin.HandlerFunc

	DeleteComment() gin.HandlerFunc

	AddCommentForComment() gin.HandlerFunc
}
