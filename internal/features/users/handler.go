package users

import "github.com/gin-gonic/gin"

type Handlers interface {
	GetUserDetails() gin.HandlerFunc
}
