package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/y-mabuchi/torasemi-todo-api/pkg/graph/generated"
	"github.com/y-mabuchi/torasemi-todo-api/pkg/graph/model"
)

func (r *queryResolver) AllUsers(ctx context.Context) ([]*model.User, error) {
	// TODO: implement
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
