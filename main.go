package main

import (
	"example/backend/rest"
	"example/backend/todo"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(rest.ErrorHandler)
	authMiddleware, err := jwt.New(rest.AuthMiddleware)
	if err != nil {
		panic(err)
	}
	r.Use(authMiddleware.MiddlewareFunc())

	todos := r.Group("/todos")
	todo.Register(todos)
	r.Run(":3001")
}
