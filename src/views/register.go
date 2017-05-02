package views

import (
	"net/http"
	"model"
)

func RegisterHandler(response http.ResponseWriter, request *http.Request)(ret int)  {
	session,collection := ConnectMongo("test", "user")
	name := request.FormValue("name")
	pass := request.FormValue("password")
	redirectTarget := "/"
	if name != "" && pass != "" {
		temp := &model.User{
			NAME: name,
			PASSWORD: pass,
		}
		collection.Insert(temp)
		session.Close()
		ret = 1
		return ret
	}
	http.Redirect(response, request, redirectTarget, 302)
}