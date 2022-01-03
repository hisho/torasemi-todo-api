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
	log.Print("action=graph.CreateUser, status=start")
	user, err := r.repo.CreateUser(ctx, input)
	if err != nil {
		log.Printf("action=r.repo.CreateUser, status=error: %v", err)
		return nil, err
	}

	return &model.User{User: user}, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input *model.UpdateUserInput) (*model.User, error) {
	log.Print("action=graph.UpdateUser, status=start")
	user, err := r.repo.UpdateUser(ctx, input)
	if err != nil {
		log.Printf("action=r.repo.UpdateUser, status=error: %v", err)
		return nil, err
	}

	return &model.User{User: user}, nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input *model.CreateTodoInput) (*model.Todo, error) {
	todo, err := r.repo.CreateTodo(ctx, input)
	if err != nil {
		log.Printf("action=graph.CreateTodo, status=error: %v", err)
		return nil, err
	}

	return &model.Todo{Todo: todo}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
