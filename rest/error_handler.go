package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		print(err)
		if re, ok := err.Err.(*RestError); ok {
			c.AbortWithStatusJSON(re.StatusCode, gin.H{"error": re.ErrorCode})
			return
		}
	}

	c.JSON(http.StatusInternalServerError, "error")
}