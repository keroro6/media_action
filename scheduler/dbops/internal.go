package dbops

import (
	"github.com/globalsign/mgo/bson"
)

func ReadVideoDeletionRecord(count int) ([]string, error) {
	session := DBS["videos"]
	sessionCopy := session.Copy()
	defer sessionCopy.Close()

	c := sessionCopy.DB("videos").C("video_del_info")

	// 使用All方法，一次性消耗较多内存，如果数据较多，可以考虑使用迭代器
	type videoDel struct {
		Id string `bson:"vid,omitempty"`
	}
	var video videoDel

	query := c.Find(bson.M{}).Batch(count).Iter()
	if query == nil {
		return nil, nil
	}

	var ids []string
	for query.Next(&video) {
		//fmt.Println(user)
		ids = append(ids, video.Id)
	}
	return ids, nil
}

func DelVideoRecord(vid string) error {

	session := DBS["videos"]
	sessionCopy := session.Copy()
	defer sessionCopy.Close()
	c := sessionCopy.DB("videos").C("video_del_info")
	return c.Remove(bson.M{"vid": vid})
}
