package dbops

import (
	"media_action/api/utils"
	"media_action/api/defs"
	"github.com/globalsign/mgo/bson"
)

var dbNameC = "comments"
var tableNameC = "comment"

func AddNewComments(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}
	session := DBS[dbNameC]
	sessionCopy := session.Copy()
	defer sessionCopy.Close() //少用defer
	c := sessionCopy.DB(dbNameC).C(tableNameC)

	comment := defs.Comment{
		Id:       id,
		AuthorId: aid,
		Content:  content,
		Ct:       defs.GetMillTime(),
		VideoId:  vid,
	}

	return c.Insert(comment)
}

func ListComments(vid string, from, to int64) ([]*defs.Comment, error) {

	session := DBS[dbNameC]
	sessionCopy := session.Copy()
	defer sessionCopy.Close() //少用defer
	c := sessionCopy.DB(dbNameC).C(tableNameC)

	var res []*defs.Comment
	q := bson.M{
		"vid": vid,
		"ct": bson.M{
			"$gt":  from,
			"$lte": to,
		},
	}

	err := c.Find(q).All(&res)
	return res, err
}
