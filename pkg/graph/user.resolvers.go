package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/y-mabuchi/torasemi-todo-api/pkg/graph/generated"
	"github.com/y-mabuchi/torasemi-todo-api/pkg/graph/model"
)

func (r *queryResolver) AllUsers(ctx context.Context) ([]*model.User, error) {
	data, err := r.repo.AllUsers(ctx)
	if err != nil {
		log.Printf("action=r.repo.AllUsers, status=error: %v", err)
		return nil, err
	}

	var users []*model.User
	for _, d := range data {
		users = append(users, &model.User{User: d})
	}

	return users, nil
}

func (r *queryResolver) User(ctx context.Context, id *int) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
