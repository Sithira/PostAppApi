package http

import (
	"RestApiBackend/internal/features/comments"
	"github.com/gin-gonic/gin"
)

func NewCommentRouter(h comments.Handler, group *gin.RouterGroup) {
	group.GET("", h.GetComments())
	group.GET(":commentId", h.GetSingleComment())
	group.POST("", h.AddComment())
	group.POST(":commentId", h.AddCommentForComment())
	group.PATCH(":commentId", h.UpdateComment())
	group.DELETE(":commentId", h.DeleteComment())
}
