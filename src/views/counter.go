package views

import (
	"gopkg.in/mgo.v2/bson"
)

type idCouter struct {
	ID int
	NAME string
}

func NewId(name string) (id int) {
	session, collection := ConnectMongo("test", "idCounter")
	collection.Upsert(bson.M{"name": name}, bson.M{"$set": bson.M{}, "$inc": bson.M{"id": 1}})
	result := idCouter{}
	collection.Find(bson.M{"name": name}).One(&result)
	session.Close()
	return result.ID
}