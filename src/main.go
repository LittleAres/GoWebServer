package main

import (
	"net/http"
	"views"
	//"fmt"
)

func main() {
	http.HandleFunc("/", views.ViewHandler)
	http.ListenAndServe(":11111", nil)
}