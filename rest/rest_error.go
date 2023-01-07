package rest

import (
	"example/backend/openapi"

	"github.com/gin-gonic/gin"
)

type RestError struct {
	StatusCode int
	ErrorCode  string
}

func (re *RestError) Error() string {
	return re.ErrorCode
}

var NotFoundError = RestError{StatusCode: 404, ErrorCode: "not_found"}
var InternalServerError = RestError{StatusCode: 500, ErrorCode: "internal_server_error"}

func AbortWithRestError(c *gin.Context, e RestError) {
	c.AbortWithStatusJSON(e.StatusCode, openapi.Error{Code: e.ErrorCode})
}

func AbortWithError(c *gin.Context, err error) {
	c.Error(err)
	c.Abort()
}
