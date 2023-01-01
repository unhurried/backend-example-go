package todo

import (
	"context"
	"example/backend/db"
	"example/backend/ent"
	"example/backend/rest"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func Register(g *gin.RouterGroup) {
	g.GET("", getList)
	g.POST("", post)
	g.GET("/:id", get)
	g.PUT("/:id", put)
	g.DELETE("/:id", del)
}

func entityToBody(e *ent.Todo) *Todo {
	return &Todo{
		Id:       strconv.Itoa(e.ID),
		Title:    e.Title,
		Category: e.Category,
		Content:  e.Content,
	}
}

func getList(c *gin.Context) {
	entities, err := db.Client.Todo.Query().All(context.Background())
	if err != nil {
		c.Error(rest.InternalServerError)
		c.Abort()
		return
	}

	items := make([]Todo, 0, len(entities))
	for _, entity := range entities {
		items = append(items, *entityToBody(entity))
	}

	var resBody = gin.H{
		"total": len(items),
		"items": items,
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, resBody)
}

func post(c *gin.Context) {
	var body Todo
	c.BindJSON(&body)

	entity, err := db.Client.Todo.Create().
		SetTitle(body.Title).SetCategory(body.Category).SetContent(body.Content).Save(context.Background())
	if err != nil {
		c.Error(rest.InternalServerError)
		c.Abort()
		return
	}
	body.Id = strconv.Itoa(entity.ID)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, body)
}

func get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	entity, err := db.Client.Todo.Get(context.Background(), id)
	if _, ok := err.(*ent.NotFoundError); ok {
		c.Error(rest.NotFoundError)
		c.Abort()
		return
	} else if err != nil {
		c.Error(rest.InternalServerError)
		c.Abort()
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, entityToBody(entity))
}

func put(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var body Todo
	c.BindJSON(&body)

	entity, err := db.Client.Todo.UpdateOneID(id).
		SetTitle(body.Title).SetCategory(body.Category).SetContent(body.Content).Save(context.Background())

	if _, ok := err.(*ent.NotFoundError); ok {
		c.Error(rest.NotFoundError)
		c.Abort()
		return
	} else if err != nil {
		c.Error(rest.InternalServerError)
		c.Abort()
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, entityToBody(entity))
}

func del(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := db.Client.Todo.DeleteOneID(id).Exec(context.Background())
	if _, ok := err.(*ent.NotFoundError); ok {
		c.Error(rest.NotFoundError)
		c.Abort()
		return
	} else if err != nil {
		c.Error(rest.InternalServerError)
		c.Abort()
		return
	}

	c.Status(http.StatusNoContent)
}
