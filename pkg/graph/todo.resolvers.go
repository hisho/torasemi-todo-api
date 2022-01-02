package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/y-mabuchi/torasemi-todo-api/pkg/graph/generated"
	"github.com/y-mabuchi/torasemi-todo-api/pkg/graph/model"
)

func (r *todoResolver) Todo(ctx context.Context, obj *model.Todo) (string, error) {
	return obj.Todo.Todo, nil
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
