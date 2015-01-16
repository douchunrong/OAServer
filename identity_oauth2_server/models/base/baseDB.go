package base

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

func Conn() *mgo.Session {
	if session == nil {
		initDBSession()
	}
	return session.Copy()
}

/*
func Close() {
	session.Close()
}
*/

func init() {
	initDBSession()
}

func initDBSession() {
	beego.Debug("111111")
	url := beego.AppConfig.String("mongodb::url")

	sess, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	session = sess
	session.SetMode(mgo.Monotonic, true)
}
