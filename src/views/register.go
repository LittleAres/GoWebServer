package views

import (
	"net/http"
	"models"
	"gopkg.in/mgo.v2/bson"
	"controllers"
)

func RegisterView(response http.ResponseWriter, request *http.Request){
	session,collection := controllers.ConnectMongo("test", "user")
	// 搞一个id
	email := request.FormValue("email")
	pass := request.FormValue("password")
	// 有没有注册过
	res := collection.Find(bson.M{"email": email}).One(&models.User{})
	if res != nil {
		collection.Insert(&models.User{
			ID:controllers.NewId("user"),
			EMAIL: email,
			PASSWORD: controllers.EncryptPass(pass),
		})
		response.Write([]byte("1"))
	} else {
		response.Write([]byte("2"))
	}
	session.Close()
}