package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	route := chi.NewRouter()
	route.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found...!", http.StatusNotFound)
	})
	fmt.Println("Stating the server on :8080 port ... !")
	http.ListenAndServe(":8080", route)
}
