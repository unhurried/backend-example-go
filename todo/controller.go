package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

var todos map[string]Todo = make(map[string]Todo)

func Register(g *gin.RouterGroup) {
	g.GET("/", getList)
	g.POST("/", post)
	g.GET("/:id", get)
	g.PUT("/:id", put)
	g.DELETE("/:id", del)
}

func getList(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, todos)
}

func post(c *gin.Context) {
	var body Todo
	c.BindJSON(&body)
	body.Id = xid.New().String()
	todos[body.Id] = body
	c.JSON(http.StatusOK, body)
}

func get(c *gin.Context) {
	id := c.Param("id")
	if todo, exists := todos[id]; exists {
		c.JSON(http.StatusOK, todo)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"code": "not_found",
		})
	}
}

func put(c *gin.Context) {
	id := c.Param("id")
	if _, exists := todos[id]; exists {
		var body Todo
		c.BindJSON(&body)
		todos[id] = body
		c.JSON(http.StatusOK, todos[id])
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"code": "not_found",
		})
	}
}

func del(c *gin.Context) {
	id := c.Param("id")
	if _, exists := todos[id]; exists {
		delete(todos, id)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"code": "not_found",
		})
	}
}
