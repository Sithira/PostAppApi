package posts

import "github.com/gin-gonic/gin"

type Handler interface {
	GetPostById() gin.HandlerFunc
	GetPostsForUser() gin.HandlerFunc
	CreatePostForUser() gin.HandlerFunc
	UpdatePostForUser() gin.HandlerFunc
	DeletePostForUser() gin.HandlerFunc
	AttachPostToUser() gin.HandlerFunc
	DeAttachPostFromUser() gin.HandlerFunc
}
