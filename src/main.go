package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"views"
)

// server main method

var router = mux.NewRouter()

func main() {
	// 编译 coffee 文件
	_ = views.StaticHandler(".coffee")
	http.Handle("/javascript/", http.StripPrefix("/javascript/",
		http.FileServer(http.Dir("src/html/javascript"))))
	http.Handle("/css/", http.StripPrefix("/css/",
		http.FileServer(http.Dir("src/html/css"))))
	http.Handle("/templates/", http.StripPrefix("/templates/",
		http.FileServer(http.Dir("src/html/templates"))))
	router.HandleFunc("/", views.PagesHandler)
	router.HandleFunc("/register_handler", views.RegisterHandler).Methods("POST")
	router.HandleFunc("/login_handler", views.LoginHandler).Methods("POST")
	router.HandleFunc("/logout_handler", views.LogoutHandler).Methods("POST")
	http.Handle("/", router)
	http.ListenAndServe(":9999", nil)
}