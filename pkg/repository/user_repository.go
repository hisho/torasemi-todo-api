package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/y-mabuchi/torasemi-todo-api/pkg/graph/model"

	"github.com/y-mabuchi/torasemi-todo-api/pkg/repository/rds"
)

func (r Repository) AllUsers(ctx context.Context) (rds.UserSlice, error) {
	log.Print("action=repository.AllUsers, status=start")
	users, err := rds.Users().All(ctx, r.db)
	if err != nil {
		log.Printf("action=AllUsers, status=error: %v", err)
		return nil, err
	}

	return users, nil
}

func (r Repository) User(ctx context.Context, id int) (*rds.User, error) {
	log.Print("action=repository.User, status=start")
	user, err := rds.FindUser(ctx, r.db, id)
	if err != nil {
		log.Printf("action=rds.Users, status=error: %v", err)
		return nil, err
	}

	return user, nil
}

func (r Repository) CreateUser(ctx context.Context, input *model.CreateUserInput) (*rds.User, error) {
	log.Print("action=repository.CreateUser, status=start")
	user := &rds.User{
		Name:      input.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := user.Insert(ctx, r.db, boil.Infer()); err != nil {
		log.Printf("action=user.Insert, status=error: %v", err)
		return nil, err
	}

	return user, nil
}

func (r Repository) UpdateUser(ctx context.Context, input *model.UpdateUserInput) (*rds.User, error) {
	log.Print("action=repository.UpdateUser, status=start")
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("action=r.db.BeginTx, statis=error: %v", err)
		return nil, err
	}

	user, err := rds.FindUser(ctx, r.db, input.ID)
	if err != nil {
		log.Printf("action=rds.FindUser, status=error: %v", err)
		tx.Rollback()
		return nil, err
	}

	user.Name = input.Name
	user.UpdatedAt = time.Now()

	_, err = user.Update(ctx, r.db, boil.Infer())
	if err != nil {
		log.Printf("action=user.Update, status=error: %v", err)
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return user, nil
}

func (r Repository) DeleteUser(ctx context.Context, input *model.DeleteUserInput) (*rds.User, error) {
	log.Print("action=repository.DeleteUser, status=start")
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("action=repository.r.db.BeginTx, status=error: %v", err)
		return nil, err
	}

	user, err := rds.FindUser(ctx, r.db, input.ID)
	if err != nil {
		log.Printf("action=repository.rds.FindUser, status=error: %v", err)
		tx.Rollback()
		return nil, err
	}

	_, err = user.Delete(ctx, r.db)
	if err != nil {
		log.Printf("action=repository.user.Delete, status=error: %v", err)
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return user, nil
}

func (r Repository) Users(ctx context.Context, input *model.UserFilter) (rds.UserSlice, error) {
	log.Print("action=repository.Users, status=start")
	var mods []qm.QueryMod

	if input != nil {
		clause := fmt.Sprintf("%s LIKE ?", rds.UserColumns.Name)
		args := "%" + *input.Name + "%"
		mods = append(mods, qm.Where(clause, args))
	}

	users, err := rds.Users(mods...).All(ctx, r.db)
	if err != nil {
		log.Printf("action=repository.rds.Users, status=error: %v", err)
		return nil, err
	}

	return users, nil
}
