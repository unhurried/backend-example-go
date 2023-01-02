package logger

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Logger *zap.Logger

func init() {
	var err error

	if gin.IsDebugging() {
		Logger, err = zap.NewDevelopment()
	} else {
		Logger, err = zap.NewProduction()
	}

	if err != nil {
		panic(err)
	}
}
