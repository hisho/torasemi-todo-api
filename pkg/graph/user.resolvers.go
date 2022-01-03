package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

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

func (r *mutationResolver) DeleteUser(ctx context.Context, input *model.DeleteUserInput) (*model.User, error) {
	log.Print("action=graph.DeleteUser, status=start")
	user, err := r.repo.DeleteUser(ctx, input)
	if err != nil {
		log.Printf("action=graph.r.repo.DeleteUser, status=error: %v", err)
		return nil, err
	}

	return &model.User{User: user}, nil
}

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

func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	user, err := r.repo.User(ctx, id)
	if err != nil {
		log.Printf("action=r.repo.User, status=error: %v", err)
		return nil, err
	}

	return &model.User{User: user}, err
}
