package controllers

import (
	"config"
	"gopkg.in/mgo.v2"
)

func ConnectMongo(dbname string, docname string)(session *mgo.Session,collection *mgo.Collection){
	// 获取当前路径
	session, _ = mgo.Dial(config.MongoUrl)
	session.SetMode(mgo.Monotonic, true)
	db := session.DB(dbname)	 //数据库名称
	collection = db.C(docname)
	return session,collection
}
