package views

import (
	"net/http"
	"model"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

func RegisterHandler(response http.ResponseWriter, request *http.Request){
	session,collection := ConnectMongo("test", "user")
	// 搞一个id
	email := request.FormValue("email")
	pass := request.FormValue("password")
	fmt.Println(email)
	// 有没有注册过
	res := collection.Find(bson.M{"email": email}).One(&model.User{})
	if res != nil {
		collection.Insert(&model.User{
			ID:NewId("user"),
			EMAIL: email,
			PASSWORD: pass,
		})
		response.Write([]byte("1"))
	} else {
		response.Write([]byte("2"))
	}
	session.Close()
}