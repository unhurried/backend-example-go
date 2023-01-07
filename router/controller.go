package router

import (
	"context"
	"example/backend/db"
	"example/backend/ent"
	"example/backend/middleware"
	"example/backend/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func entityToBody(e *ent.Todo) *model.Todo {
	return &model.Todo{
		Id:       strconv.Itoa(e.ID),
		Title:    e.Title,
		Category: e.Category,
		Content:  e.Content,
	}
}

func GetList(c *gin.Context) {
	entities, err := db.Client.Todo.Query().All(context.Background())
	if err != nil {
		middleware.AbortWithError(c, err)
		return
	}

	items := make([]model.Todo, 0, len(entities))
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

func Post(c *gin.Context) {
	var body model.Todo
	c.BindJSON(&body)

	entity, err := db.Client.Todo.Create().
		SetTitle(body.Title).SetCategory(body.Category).SetContent(body.Content).Save(context.Background())
	if err != nil {
		middleware.AbortWithError(c, err)
		return
	}
	body.Id = strconv.Itoa(entity.ID)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, body)
}

func Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	entity, err := db.Client.Todo.Get(context.Background(), id)
	if _, ok := err.(*ent.NotFoundError); ok {
		middleware.AbortWithRestError(c, middleware.NotFoundError)
		return
	} else if err != nil {
		middleware.AbortWithError(c, err)
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, entityToBody(entity))
}

func Put(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var body model.Todo
	c.BindJSON(&body)

	entity, err := db.Client.Todo.UpdateOneID(id).
		SetTitle(body.Title).SetCategory(body.Category).SetContent(body.Content).Save(context.Background())

	if _, ok := err.(*ent.NotFoundError); ok {
		middleware.AbortWithRestError(c, middleware.NotFoundError)
		return
	} else if err != nil {
		middleware.AbortWithError(c, err)
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, entityToBody(entity))
}

func Del(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := db.Client.Todo.DeleteOneID(id).Exec(context.Background())
	if _, ok := err.(*ent.NotFoundError); ok {
		middleware.AbortWithRestError(c, middleware.NotFoundError)
		return
	} else if err != nil {
		middleware.AbortWithError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
