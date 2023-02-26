package rest

import (
	"context"
	"example/backend/db"
	"example/backend/ent"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type Server struct{}

func (s Server) TodoControllerGetList(ctx echo.Context, params TodoControllerGetListParams) error {
	entities, err := db.Client.Todo.Query().All(context.Background())
	if err != nil {
		return err
	}

	items := make([]Todo, 0, len(entities))
	for _, entity := range entities {
		items = append(items, *entityToBody(entity))
	}

	total := len(items)
	var resBody = TodoList{
		Total: &total,
		Items: &items,
	}

	ctx.Response().Header().Set("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, resBody)
	return nil
}

func (s Server) TodoControllerPost(ctx echo.Context) error {
	var body Todo
	ctx.Bind(&body)

	content := ""
	if body.Content != nil {
		content = *body.Content
	}

	entity, err := db.Client.Todo.Create().
		SetTitle(body.Title).
		SetCategory(string(body.Category)).
		SetContent(content).Save(context.Background())
	if err != nil {
		return err
	}

	ctx.Response().Header().Set("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, *entityToBody(entity))
	return nil
}

func (s Server) TodoControllerDelete(ctx echo.Context, id string) error {
	idAsInt, _ := strconv.Atoi(id)

	err := db.Client.Todo.DeleteOneID(idAsInt).Exec(context.Background())
	if _, ok := err.(*ent.NotFoundError); ok {
		return &NotFoundError
	} else if err != nil {
		return err
	}

	ctx.NoContent(http.StatusNoContent)
	return nil
}

func (s Server) TodoControllerGet(ctx echo.Context, id string) error {
	idAsInt, _ := strconv.Atoi(id)

	entity, err := db.Client.Todo.Get(context.Background(), idAsInt)
	if _, ok := err.(*ent.NotFoundError); ok {
		return &NotFoundError
	} else if err != nil {
		return err
	}

	ctx.Response().Header().Set("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, *entityToBody(entity))
	return nil
}

func (s Server) TodoControllerUpdate(ctx echo.Context, id string) error {
	idAsInt, _ := strconv.Atoi(id)
	var body Todo
	ctx.Bind(&body)

	query := db.Client.Todo.UpdateOneID(idAsInt).
		SetTitle(body.Title).
		SetCategory(string(body.Category))

	if body.Content != nil {
		query.SetContent(*body.Content)
	}

	entity, err := query.Save(context.Background())
	if _, ok := err.(*ent.NotFoundError); ok {
		return &NotFoundError
	} else if err != nil {
		return err
	}

	ctx.Response().Header().Set("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, *entityToBody(entity))
	return nil
}

func entityToBody(e *ent.Todo) *Todo {
	id := strconv.Itoa(e.ID)
	return &Todo{
		Id:       &id,
		Title:    e.Title,
		Category: TodoCategory(e.Category),
		Content:  &e.Content,
	}
}
