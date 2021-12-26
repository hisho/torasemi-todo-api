package main

import (
	"log"

	"github.com/y-mabuchi/torasemi-todo-api/pkg/config"
	"github.com/y-mabuchi/torasemi-todo-api/pkg/db/migration"
	"github.com/y-mabuchi/torasemi-todo-api/pkg/infrastructure"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	conf := config.NewConfigFromEnv()
	db := infrastructure.NewDB(conf.DBConfig)
	defer db.Close()

	if err := migration.Run(db); err != nil {
		log.Fatalf("action=migration.Run, status=error: %v", err)
	}
}
