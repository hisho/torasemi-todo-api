package model

import "github.com/y-mabuchi/torasemi-todo-api/pkg/repository/rds"

type Todo struct {
	*rds.Todo
}
