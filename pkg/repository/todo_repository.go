package repository

import (
	"context"
	"log"

	"github.com/y-mabuchi/torasemi-todo-api/pkg/repository/rds"
)

func (r Repository) AllTodos(ctx context.Context) (rds.TodoSlice, error) {
	todos, err := rds.Todos().All(ctx, r.db)
	if err != nil {
		log.Printf("action=rds.Todos, status=error: %v", err)
		return nil, err
	}

	return todos, nil
}
