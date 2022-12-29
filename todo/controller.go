package todo

import (
	"example/backend/rest"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

var todos map[string]Todo = make(map[string]Todo)

func Register(g *gin.RouterGroup) {
	g.GET("", getList)
	g.POST("", post)
	g.GET("/:id", get)
	g.PUT("/:id", put)
	g.DELETE("/:id", del)
}

func getList(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	items := make([]Todo, 0, len(todos))
	for _, val := range todos {
		items = append(items, val)
	}

	var resBody = gin.H{
		"total": len(todos),
		"items": items,
	}
	c.JSON(http.StatusOK, resBody)
}

func post(c *gin.Context) {
	var body Todo
	c.BindJSON(&body)
	body.Id = xid.New().String()
	todos[body.Id] = body
	c.JSON(http.StatusCreated, body)
}

func get(c *gin.Context) {
	id := c.Param("id")
	if todo, exists := todos[id]; exists {
		c.JSON(http.StatusOK, todo)
	} else {
		c.Error(rest.NotFoundError)
		c.Abort()
		return
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
		c.Error(rest.NotFoundError)
		c.Abort()
		return
	}
}

func del(c *gin.Context) {
	id := c.Param("id")
	if _, exists := todos[id]; exists {
		delete(todos, id)
		c.Status(http.StatusNoContent)
	} else {
		c.Error(rest.NotFoundError)
		c.Abort()
		return
	}
}
