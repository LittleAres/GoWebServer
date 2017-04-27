package views

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"fmt"
)

var url_prefix = "src/html/templates/"

type Page struct {
	Title string
	Body  []byte
}

func LoadPage(title string) (*Page, error) {
	filename := url_prefix + title + ".html"
	fmt.Println("File name: ",filename)
	body, _ := ioutil.ReadFile(filename)
	//if err != nil {
	//	filename := url_prefix + "404.html"
	//	body, _ := ioutil.ReadFile(filename)
	//	return &Page{Title: "404", Body: body}, err
	//}
	return &Page{Title: title, Body: body}, nil
}

func RenderTemplate(w http.ResponseWriter, title string, p *Page) {
	t, _ := template.ParseFiles(url_prefix + title + ".html")
	fmt.Println(t)
	t.Execute(w, p)
}

func PagesHandler(response http.ResponseWriter, request *http.Request) {
	userName := getUserName(request)
	var title = request.URL.Path[len("/"):]
	if userName != "" {
		if title == "" {
			title = "index"
		}
	} else {
		title = "login"
	}
	p, _ := LoadPage(title)
	//if err != nil {
	//	RenderTemplate(response, "404", p)
	//} else {
	RenderTemplate(response, title, p)
	//}
}