package auth

import "github.com/gin-gonic/gin"

type AuthenticationHandler interface {
	Login() gin.HandlerFunc
}
