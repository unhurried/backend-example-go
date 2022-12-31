package db

import (
	"context"
	"example/backend/ent"
	"example/backend/env"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var Client *ent.Client

func Open() {
	c, err := ent.Open(env.CONFIG.DB_DRIVER, env.CONFIG.DB_DATA_SOURCE)
	if err != nil {
		panic(fmt.Sprintf("failed opening connection to sqlite: %v", err))
	}

	if err := c.Schema.Create(context.Background()); err != nil {
		panic(fmt.Sprintf("failed creating schema resources: %v", err))
	}

	Client = c
}

func Close() {
	Client.Close()
}
