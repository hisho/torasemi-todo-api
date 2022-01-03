package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/y-mabuchi/torasemi-todo-api/pkg/graph/generated"
	"github.com/y-mabuchi/torasemi-todo-api/pkg/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input *model.CreateTodoInput) (*model.Todo, error) {
	todo, err := r.repo.CreateTodo(ctx, input)
	if err != nil {
		log.Printf("action=graph.CreateTodo, status=error: %v", err)
		return nil, err
	}

	return &model.Todo{Todo: todo}, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input *model.UpdateTodoInput) (*model.Todo, error) {
	log.Print("action=graph.UpdateTodo, status=start")
	todo, err := r.repo.UpdateTodo(ctx, input)
	if err != nil {
		log.Printf("action=graph.r.repo.UpdateTodo, status=error: %v", err)
		return nil, err
	}

	return &model.Todo{Todo: todo}, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input *model.DeleteTodoInput) (*model.Todo, error) {
	log.Print("action=graph.DeleteTodo, status=start")
	todo, err := r.repo.DeleteTodo(ctx, input)
	if err != nil {
		log.Printf("action=graph.r.repo.DeleteTodo, status=error: %v", err)
		return nil, err
	}

	return &model.Todo{Todo: todo}, nil
}

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

func (r *queryResolver) Todo(ctx context.Context, id int) (*model.Todo, error) {
	todo, err := r.repo.Todo(ctx, id)
	if err != nil {
		log.Printf("action=r.repo.Todo, status=error: %v", err)
		return nil, err
	}

	return &model.Todo{Todo: todo}, nil
}

func (r *queryResolver) Todos(ctx context.Context, filter *model.TodoFilter) ([]*model.Todo, error) {
	log.Print("action=graph.Todos, status=start")
	data, err := r.repo.Todos(ctx, filter)
	if err != nil {
		log.Printf("action=graph.Todos, status=error: %v", err)
		return nil, err
	}

	var todos []*model.Todo
	for _, d := range data {
		todos = append(todos, &model.Todo{
			Todo: d,
		})
	}

	return todos, nil
}

func (r *todoResolver) Todo(ctx context.Context, obj *model.Todo) (string, error) {
	return obj.Todo.Todo, nil
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
