package internal

import (
	"fmt"
	"todox-with-sqlite/internal/todos"

	"github.com/leapkit/core/server"
)

func AddServices(r server.Router) error {
	db, err := DB()
	if err != nil {
		return fmt.Errorf("connecting to the database: %w", err)
	}

	r.Use(server.InCtxMiddleware("todos", todos.NewService(db)))
	return nil
}
