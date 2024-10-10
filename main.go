package main

import (
	"log"
	"net/http"

	"github.com/webdevfuel/htmx-loading-indicator/template"
	"github.com/webdevfuel/htmx-loading-indicator/user"
)

func main() {
	router := http.NewServeMux()
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	router.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		u := user.ListUsers()
		component := template.Users(u)
		component.Render(r.Context(), w)
	})
	router.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		e := r.FormValue("email")
		u, err := user.AddUser(e)
		if err != nil {
			log.Printf("error adding user %v", err)
			return
		}
		component := template.UserRow(u)
		component.Render(r.Context(), w)
	})
	router.HandleFunc("DELETE /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		err := user.DeleteUser(id)
		if err != nil {
			log.Printf("error deleting user %v", err)
			return
		}
	})
	http.ListenAndServe("localhost:3000", router)
}
