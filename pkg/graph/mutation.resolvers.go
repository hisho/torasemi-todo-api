package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/y-mabuchi/torasemi-todo-api/pkg/graph/generated"
	"github.com/y-mabuchi/torasemi-todo-api/pkg/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.CreateUserInput) (*model.User, error) {
	user, err := r.repo.CreateUser(ctx, input)
	if err != nil {
		log.Printf("action=r.repo.CreateUser, status=error: %v", err)
		return nil, err
	}

	return &model.User{User: user}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
