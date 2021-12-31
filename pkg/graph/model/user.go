package model

import "github.com/y-mabuchi/torasemi-todo-api/pkg/repository/rds"

type User struct {
	*rds.User
}
