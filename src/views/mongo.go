package views

import (
	"gopkg.in/mgo.v2"
)

const (
	URL = "127.0.0.1:27017"
)

func ConnectMongo(dbname string, docname string)(session *mgo.Session,collection *mgo.Collection){
	// 获取当前路径
	session, _ = mgo.Dial(URL)
	session.SetMode(mgo.Monotonic, true)
	db := session.DB(dbname)	 //数据库名称
	collection = db.C(docname)
	return session,collection
}
