package views

import (
	"controllers"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"models"
)

// login view
func LoginView(response http.ResponseWriter, request *http.Request){
	session,collection := controllers.ConnectMongo("test", "user")
	// 搞一个id
	email := request.FormValue("email")
	pass := request.FormValue("password")
	// 有没有注册过
	result := &models.User{}
	err := collection.Find(bson.M{"email": email}).One(result)
	if err == nil {
		if result.PASSWORD == controllers.EncryptPass(pass) {
			controllers.SetSession(email, response)
			response.Write([]byte("1"))
		} else {
			response.Write([]byte("2"))
		}
	} else {
		response.Write([]byte("3"))
	}
	session.Close()
}

// logout view
func LogoutView(response http.ResponseWriter, request *http.Request) {
	controllers.ClearSession(response)
	http.Redirect(response, request, "/", 302)
}