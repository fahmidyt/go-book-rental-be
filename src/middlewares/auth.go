package middlewares

import (
	"github.com/fahmidyt/go-book-rental-be/src/helpers"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return helpers.VerifyAccesToken
}

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: add auth token validation
		c.Next()
	}
}
