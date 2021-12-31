package repository

import (
	"context"
	"log"

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
