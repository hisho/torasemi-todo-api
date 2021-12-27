package server

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/go-chi/chi"

	"github.com/y-mabuchi/torasemi-todo-api/pkg/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"

	"github.com/y-mabuchi/torasemi-todo-api/pkg/graph"
	"github.com/y-mabuchi/torasemi-todo-api/pkg/repository"

	_ "github.com/go-sql-driver/mysql"
	"github.com/y-mabuchi/torasemi-todo-api/pkg/config"
	"github.com/y-mabuchi/torasemi-todo-api/pkg/infrastructure"
)

func Run() error {
	router := chi.NewRouter()

	conf := config.NewConfigFromEnv()
	db := infrastructure.NewDB(conf.DBConfig)
	defer db.Close()

	repo := repository.NewRepository(db)
	resolver := graph.NewResolver(repo)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	router.Handle("/", playground.Handler("GraphQL Playground", "/api/query"))
	router.Handle("/api/query", srv)

	log.Print("listen http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}

	return nil
}
