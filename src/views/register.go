package views

import (

)
import (
	//"gopkg.in/mgo.v2/bson"
	//"fmt"
	"net/http"
)

type User struct {
	NAME  string
	PASSWORD string
}

func RegisterHandler(response http.ResponseWriter, request *http.Request)  {
	session,collection := ConnectMongo("test", "user")
	name := request.FormValue("name")
	pass := request.FormValue("password")
	redirectTarget := "/login"
	if name != "" && pass != "" {
		temp := &User{
			NAME: name,
			PASSWORD: pass,
		}
		collection.Insert(temp)
		session.Close()
	}
	http.Redirect(response, request, redirectTarget, 302)
}