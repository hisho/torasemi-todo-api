package main

import (
	"log"

	"github.com/y-mabuchi/torasemi-todo-api/pkg/server"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
