package main

import (
	"assignment/cmd/database"
	"assignment/controllers"
	"assignment/views"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func main() {
	/* Requiring Database */
	db := database.Connect()
	defer db.Close()

	/* _, err := db.Exec(`

	CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT UNIQUE NOT NULL
	)`)

	if err != nil {
		panic(err)
	} */

	route := chi.NewRouter()

	route.Get("/", controllers.StaticHandler(views.Must(views.Parse(filepath.Join("templates", "dashboard.gohtml")))))

	route.Get("/login", controllers.StaticHandler(views.Must(views.Parse(filepath.Join("templates", "login.gohtml")))))

	route.Get("/register", controllers.StaticHandler(views.Must(views.Parse(filepath.Join("templates", "register.gohtml")))))

	route.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found...!", http.StatusNotFound)
	})
	fmt.Println("Stating the server on :8080 port ... !")
	http.ListenAndServe(":8080", route)
}
