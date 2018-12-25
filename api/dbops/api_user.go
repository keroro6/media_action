package dbops

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"media_action/api/defs"
)

func AddUserCredential(loginName string, pwd string) error {
	session := DBS["users"]
	sessionCopy := session.Copy()
	defer sessionCopy.Close()

	c := sessionCopy.DB("users").C("user")

	user := &defs.UserCredential{
		UserName: loginName,
		Pwd:      pwd,
		Ct:       defs.GetMillTime(),
	}
	return c.Insert(user)
}
func GetUserCredential(loginName string) (string, error) {
	session := DBS["users"]
	sessionCopy := session.Copy()
	defer sessionCopy.Close() //少用defer

	c := sessionCopy.DB("users").C("user")

	var user *defs.UserCredential
	err := c.Find(bson.M{"user_name": loginName}).One(&user)
	if err != nil {
		if err == mgo.ErrNotFound {
			err = nil
		}
		return "", err
	}
	return user.Pwd, nil
}
func DeleteUser(loginName string, pwd string) error {
	session := DBS["users"]
	sessionCopy := session.Copy()
	defer sessionCopy.Close() //少用defer

	c := sessionCopy.DB("users").C("user")

	user := &defs.UserCredential{
		UserName: loginName,
		Pwd:      pwd,
	}
	return c.Remove(user)
}
