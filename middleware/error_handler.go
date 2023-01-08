package middleware

import (
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) != 0 {
		AbortWithRestError(c, InternalServerError)
	}
}
