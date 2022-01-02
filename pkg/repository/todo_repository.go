package repository

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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