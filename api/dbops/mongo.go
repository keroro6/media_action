package dbops

import (
	"github.com/astaxie/beego/logs"
	"github.com/globalsign/mgo"
)

var (
	// DBS 表示mongo连接，key是db名，value是db连接
	DBS map[string]*mgo.Session = map[string]*mgo.Session{}
)
var tables = map[string]string{
	"users":    "127.0.0.1:27017",
	"videos":   "127.0.0.1:27017",
	"comments": "127.0.0.1:27017",
	"sessions": "127.0.0.1:27017",
}

func init() {
	for key, value := range tables {
		session, err := mgo.Dial(value)
		if err != nil {
			logs.Error("open mgo con err:%v, key:%s", err, key)
			panic(err.Error())
			return
		}
		DBS[key] = session
	}
}
