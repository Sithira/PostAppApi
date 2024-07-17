package http

import (
	"RestApiBackend/internal/features/posts"
	"github.com/gin-gonic/gin"
)

func NewPostRouter(h posts.Handler, group *gin.RouterGroup) {
	group.GET("", h.GetPostsForUser())
	group.POST("", h.CreatePostForUser())
	group.GET(":postId", h.GetPostById())
	group.PUT(":postId", h.UpdatePostForUser())
	group.DELETE(":postId", h.DeletePostForUser())
}
