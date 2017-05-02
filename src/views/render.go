package views

import (
	"html/template"
	"io/ioutil"
	"net/http"
	//"fmt"
	"fmt"
)

var url_prefix = "src/html/templates/"

type Page struct {
	Title string
	Body  []byte
}

func LoadPage(title string) (*Page, error) {
	filename := url_prefix + title + ".html"
	fmt.Println(filename)
	body, _ := ioutil.ReadFile(filename)
	return &Page{Title: title, Body: body}, nil
}

func RenderTemplate(w http.ResponseWriter, title string, p *Page) {
	t, _ := template.ParseFiles(url_prefix + title + ".html")
	t.Execute(w, p)
}

func PagesHandler(response http.ResponseWriter, request *http.Request) {
	userName := getUserName(request)
	var title = request.URL.Path[len("/"):]
	if userName != "" {
		if title == "" {
			title = "base"
		}
	} else {
		title = "login"
	}
	p, _ := LoadPage(title)
	RenderTemplate(response, title, p)
}