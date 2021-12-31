package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/y-mabuchi/torasemi-todo-api/pkg/graph/model"
)

func (r *queryResolver) AllTodos(ctx context.Context) ([]*model.Todo, error) {
	data, err := r.repo.AllTodos(ctx)
	if err != nil {
		log.Printf("action=r.repo.AllTodos, status=error: %v", err)
		return nil, err
	}
	var todos []*model.Todo
	for _, d := range data {
		todos = append(todos, &model.Todo{Todo: d})
	}

	return todos, nil
}
