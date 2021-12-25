package server

import (
	"fmt"
	"log"
	"net/http"
)

func Run() error {
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
