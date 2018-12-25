package dbops

import "github.com/globalsign/mgo/bson"

func AddVideoDeletionRecord(vid string) error {
	session := DBS["videos"]
	sessionCopy := session.Copy()
	defer sessionCopy.Close()

	c := sessionCopy.DB("videos").C("video_del_info")

	doc := bson.M{
		"vid": vid,
	}
	return c.Insert(&doc)

}
