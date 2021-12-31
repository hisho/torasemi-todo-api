package repository

import (
	"context"
	"log"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/y-mabuchi/torasemi-todo-api/pkg/graph/model"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/y-mabuchi/torasemi-todo-api/pkg/repository/rds"
)

func (r Repository) AllUsers(ctx context.Context) (rds.UserSlice, error) {
	users, err := rds.Users().All(ctx, r.db)
	if err != nil {
		log.Printf("action=AllUsers, status=error: %v", err)
		return nil, err
	}

	return users, nil
}

func (r Repository) User(ctx context.Context, id int) (*rds.User, error) {
	var mods []qm.QueryMod
	mods = append(mods, rds.UserWhere.ID.EQ(id))
	user, err := rds.Users(mods...).One(ctx, r.db)
	if err != nil {
		log.Printf("action=rds.Users, status=error: %v", err)
		return nil, err
	}
	return user, nil
}

func (r Repository) CreateUser(ctx context.Context, input *model.CreateUserInput) (*rds.User, error) {
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
