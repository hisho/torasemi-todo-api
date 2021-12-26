package server

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/y-mabuchi/torasemi-todo-api/pkg/config"
	"github.com/y-mabuchi/torasemi-todo-api/pkg/infrastructure"
)

func Run() error {
	conf := config.NewConfigFromEnv()
	db := infrastructure.NewDB(conf.DBConfig)
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Printf("action=db.Ping, status=error: %v", err)
	}

	http.HandleFunc("/", helloHandler)

	log.Print("listen http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	return nil
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
