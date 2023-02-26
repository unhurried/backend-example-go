package logger

import (
	"example/backend/env"

	"go.uber.org/zap"
)

var Logger *zap.Logger

func init() {
	var err error

	if env.CONFIG.DEBUG_MODE == "true" {
		Logger, err = zap.NewDevelopment()
	} else {
		Logger, err = zap.NewProduction()
	}

	if err != nil {
		panic(err)
	}
}
