package views

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"fmt"
	//"reflect"
)

var url_prefix = "src/htmls/templates/"

type Page struct {
	Title string
	Body  []byte
}

func loadPage(title string) (*Page, error) {
	filename := url_prefix + title + ".html"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		filename := url_prefix + "404.html"
		body, _ := ioutil.ReadFile(filename)
		return &Page{Title: "404", Body: body}, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, title string, p *Page) {
	t, _ := template.ParseFiles(url_prefix + title + ".html")
	t.Execute(w, p)
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]
	fmt.Println(r.URL.Path)
	fmt.Println("title: "+title)
	p, err := loadPage(title)
	if err != nil {
		renderTemplate(w, "404", p)
	} else {
		renderTemplate(w, title, p)
	}
}