package main

import (
	"example/backend/todo"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	todos := r.Group("/todos")
	todo.Register(todos)
	r.Run()
}
