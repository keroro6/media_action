package dbops

import (
	"testing"
	"fmt"
)

func clearTables() {
	session := DBS["users"]
	sessionCopy := session.Copy()
	c := sessionCopy.DB("users").C("user")
	err := c.DropCollection()
	if err != nil {
		fmt.Printf("clear tables err:%v\n", err)
		return
	}
	sessionCopy.Close()

	session = DBS["videos"]
	sessionCopy = session.Copy()
	c = sessionCopy.DB("videos").C("video")
	err = c.DropCollection()
	if err != nil {
		fmt.Printf("clear tables err:%v\n", err)
		return
	}
	sessionCopy.Close()

	session = DBS["comments"]
	sessionCopy = session.Copy()
	c = sessionCopy.DB("comments").C("comment")
	err = c.DropCollection()
	if err != nil {
		fmt.Printf("clear tables err:%v\n", err)
		return
	}
	sessionCopy.Close()

	session = DBS["sessions"]
	sessionCopy = session.Copy()
	c = sessionCopy.DB("sessions").C("session")
	err = c.DropCollection()
	if err != nil {
		fmt.Printf("clear tables err:%v\n", err)
		return
	}
	sessionCopy.Close()
}
func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testRegetUser)

}

func testAddUser(t *testing.T) {

	err := AddUserCredential("wgy", "123")
	if err != nil {
		t.Errorf("error of add user :%v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("wgy")
	if err != nil || pwd != "123" {
		t.Errorf("error of get user :%v", err)
	}
	t.Log(pwd)
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("wgy", "123")
	if err != nil {
		t.Errorf("error of delete user :%v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("wgy")
	if err != nil || pwd != "" {
		t.Errorf("error of get user :%v", err)
	}
}
