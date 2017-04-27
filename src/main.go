package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"views"
)

// server main method

var router = mux.NewRouter()

func main() {
	templates := views.GetTemplates()
	for i := range templates {
		router.HandleFunc("/" + templates[i], views.PagesHandler)
	}
	router.HandleFunc("/", views.PagesHandler)
	router.HandleFunc("/login_handler", views.LoginHandler).Methods("POST")
	router.HandleFunc("/logout_handler", views.LogoutHandler).Methods("POST")
	http.Handle("/", router)
	http.ListenAndServe(":8000", nil)
}