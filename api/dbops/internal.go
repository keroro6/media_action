package dbops

import (
	"media_action/api/defs"
	"github.com/globalsign/mgo/bson"
	"sync"
)

var dbNameS = "sessions"
var dbTableS = "session"

func InsertSession(sid string, ttl int64, uname string) error {
	session := DBS[dbNameS]
	sessionCopy := session.Copy()
	defer sessionCopy.Close() //少用defer
	c := sessionCopy.DB(dbNameS).C(dbTableS)

	se := defs.SimpleSession{
		SessionId: sid,
		UserName:  uname,
		TTL:       ttl,
	}

	return c.Insert(se)
}

func RetrieveSession(sid string) (se *defs.SimpleSession, err error) {
	session := DBS[dbNameS]
	sessionCopy := session.Copy()
	defer sessionCopy.Close() //少用defer
	c := sessionCopy.DB(dbNameS).C(dbTableS)

	err = c.Find(bson.M{"sid": sid}).One(&se)
	return
}
func RetrieveAllSessions() (*sync.Map, error) {
	session := DBS[dbNameS]
	sessionCopy := session.Copy()
	defer sessionCopy.Close() //少用defer
	c := sessionCopy.DB(dbNameS).C(dbTableS)

	var ss []*defs.SimpleSession
	err := c.Find(bson.M{}).All(&ss)
	if err != nil {
		return nil, err
	}
	m := &sync.Map{}
	for _, s := range ss {
		m.Store(s.SessionId, s)
	}
	return m, nil
}

func DeleteSession(sid string) error {
	session := DBS[dbNameS]
	sessionCopy := session.Copy()
	defer sessionCopy.Close() //少用defer
	c := sessionCopy.DB(dbNameS).C(dbTableS)

	return c.Remove(bson.M{"sid": sid})
}
