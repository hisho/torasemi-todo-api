package repository

import (
	"context"
	"fmt"
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

func (r Repository) UpdateTodo(ctx context.Context, input *model.UpdateTodoInput) (*rds.Todo, error) {
	log.Print("action=repository.UpdateTodo, status=start")
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("action=repository.r.db.BeginTx, status=error: %v", err)
		return nil, err
	}

	todo, err := rds.FindTodo(ctx, r.db, input.ID)
	if err != nil {
		log.Printf("action=repository.UpdateTodo, status=error: %v", err)
		tx.Rollback()
		return nil, err
	}

	todo.Todo = input.Todo
	todo.Finished = input.Finished
	todo.UserID = input.UserID
	todo.UpdatedAt = time.Now()

	if _, err := todo.Update(ctx, r.db, boil.Infer()); err != nil {
		log.Printf("action=repository.todo.Update, status=error: %v", err)
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return todo, nil
}

func (r Repository) DeleteTodo(ctx context.Context, input *model.DeleteTodoInput) (*rds.Todo, error) {
	log.Print("action=repository.DeleteTodo, status=start")
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("action=repository.r.db.BeginTx, status=error: %v", err)
		return nil, err
	}

	todo, err := rds.FindTodo(ctx, r.db, input.ID)
	if err != nil {
		log.Printf("action=repository.rds.FindTodo, status=error: %v", err)
		tx.Rollback()
		return nil, err
	}

	if _, err := todo.Delete(ctx, r.db); err != nil {
		log.Printf("action=repository.todo.Delete, status=error: %v", err)
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return todo, nil
}

func (r Repository) Todos(ctx context.Context, filter *model.TodoFilter) (rds.TodoSlice, error) {
	log.Print("action=repository.Todos, status=start")
	var mods []qm.QueryMod

	if filter != nil && filter.Todo != nil {
		clause := fmt.Sprintf("%s LIKE ?", rds.TodoColumns.Todo)
		args := "%" + *filter.Todo + "%"
		mods = append(mods, qm.Where(clause, args))
	}

	if filter != nil && filter.Finished != nil {
		mods = append(mods, rds.TodoWhere.Finished.EQ(*filter.Finished))
	}

	if filter != nil && filter.UserID != nil {
		mods = append(mods, rds.TodoWhere.UserID.EQ(*filter.UserID))
	}

	todos, err := rds.Todos(mods...).All(ctx, r.db)
	if err != nil {
		log.Printf("action=repository.rds.Todos, status=error: %v", err)
		return nil, err
	}

	return todos, nil
}
