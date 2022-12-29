package main

import (
	"example/backend/rest"
	"example/backend/todo"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(rest.ErrorHandler)
	todos := r.Group("/todos")
	todo.Register(todos)
	r.Run(":3001")
}
