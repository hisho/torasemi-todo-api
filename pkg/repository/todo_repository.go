package repository

import (
	"context"
	"log"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/y-mabuchi/torasemi-todo-api/pkg/graph/model"

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

func (r Repository) Todo(ctx context.Context, id int) (*rds.Todo, error) {
	var mods []qm.QueryMod
	mods = append(mods, rds.TodoWhere.ID.EQ(id))
	todo, err := rds.Todos(mods...).One(ctx, r.db)
	if err != nil {
		log.Printf("action=rds.Todos, status=error: %v", err)
		return nil, err
	}

	return todo, nil
}

func (r Repository) CreateTodo(ctx context.Context, input *model.CreateTodoInput) (*rds.Todo, error) {
	log.Print("action=repository.CreateTodo, status=start")
	todo := &rds.Todo{
		Todo:      input.Todo,
		UserID:    input.UserID,
		Finished:  false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := todo.Insert(ctx, r.db, boil.Infer()); err != nil {
		log.Printf("action=repository.todo.Insert, status=error: %v", err)
		return nil, err
	}

	return todo, nil
}
