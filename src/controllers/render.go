package controllers

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

var url_prefix = "src/html/templates/"

type Page struct {
	Title string
	Body  []byte
}

func LoadPage(title string) (*Page, error) {
	filename := url_prefix + title + ".html"
	body, _ := ioutil.ReadFile(filename)
	return &Page{Title: title, Body: body}, nil
}

func RenderTemplate(w http.ResponseWriter, title string, p *Page) {
	t, _ := template.ParseFiles(url_prefix + title + ".html")
	t.Execute(w, p)
}