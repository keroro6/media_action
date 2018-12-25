package defs

import "time"

type UserCredential struct {
	UserName string `json:"user_name,omitempty" bson:"user_name,omitempty"`
	Pwd      string `json:"pwd,omitempty" bson:"pwd,omitempty"`
	Ct       int64  `json:"ct,omitempty" bson:"ct,omitempty"`
}

type VideoInfo struct {
	Id           string `json:"id" bson:"id"`
	AuthorId     int    `json:"aid" bson:"aid"`
	Name         string `json:"name" bson:"name"`
	DisplayCtime string `json:"dct" bson:"dct"`
	Ct           int64  `json:"ct" bson:"ct"`
}

type Comment struct {
	Id       string `json:"id" bson:"id"`
	AuthorId int    `json:"aid" bson:"aid"`
	Content  string `json:"content" bson:"content"`
	Ct       int64  `json:"ct" bson:"ct"`
	VideoId  string `json:"vid" bson:"vid"`
}
type SimpleSession struct {
	SessionId string `json:"sid" bson:"sid"`
	UserName  string `json:"user_name" bson:"user_name"` //用户登录名
	TTL       int64  `json:"ttl" bson:"ttl"`             //session过期时间
}

func GetMillTime() int64 {
	return time.Now().UnixNano() / 1000000
}

type SignedUp struct {
	Success   bool   `json:"success"`
	SessionId string `json:"sid"`
}
