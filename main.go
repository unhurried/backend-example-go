package main

import (
	"example/backend/db"
	"example/backend/logger"
	"example/backend/rest"
	"example/backend/todo"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Open()
	defer db.Close()

	r := gin.New()

	r.Use(ginzap.Ginzap(logger.Logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger.Logger, true))

	r.Use(rest.ErrorHandler)

	authMiddleware, err := jwt.New(rest.AuthMiddleware)
	if err != nil {
		logger.Logger.Error(err.Error())
		panic(err)
	}
	r.Use(authMiddleware.MiddlewareFunc())

	todo.AddRoutes(r)
	r.Run(":3001")
}
