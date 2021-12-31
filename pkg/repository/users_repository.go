package repository

import (
	"context"
	"log"

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
