package dbops

import (
	"media_action/api/defs"
	"media_action/api/utils"
	"time"
	"github.com/globalsign/mgo/bson"
)

var dbName = "videos"
var tableName = "video"

func AddNewVideo(aid int, name string) (videoInfo *defs.VideoInfo, err error) {

	session := DBS[dbName]
	sessionCopy := session.Copy()
	defer sessionCopy.Close() //少用defer
	c := sessionCopy.DB(dbName).C(tableName)

	vid, err := utils.NewUUID()
	if err != nil {
		return
	}

	ct := defs.GetMillTime()
	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")

	videoInfo = &defs.VideoInfo{
		Id:           vid,
		AuthorId:     aid,
		Name:         name,
		DisplayCtime: ctime,
		Ct:           ct,
	}
	err = c.Insert(videoInfo)
	return
}
func GetVideoInfo(vid string) (videoInfo *defs.VideoInfo, err error) {
	session := DBS[dbName]
	sessionCopy := session.Copy()
	defer sessionCopy.Close() //少用defer
	c := sessionCopy.DB(dbName).C(tableName)

	err = c.Find(bson.M{"id": vid}).One(&videoInfo)
	return
}

func DeleteVideoInfo(vid string) error {
	session := DBS[dbName]
	sessionCopy := session.Copy()
	defer sessionCopy.Close() //少用defer
	c := sessionCopy.DB(dbName).C(tableName)

	return c.Remove(bson.M{"id": vid})
}
