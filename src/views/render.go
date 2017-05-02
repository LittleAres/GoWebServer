package views

import (
	"net/http"
	"controllers"
)

func RenderView(response http.ResponseWriter, request *http.Request) {
	userName := controllers.GetUserName(request)
	var title = request.URL.Path[len("/"):]
	if userName != "" {
		if title == "" {
			title = "base"
		}
	} else {
		title = "login"
	}
	p, _ := controllers.LoadPage(title)
	controllers.RenderTemplate(response, title, p)
}