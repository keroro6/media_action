package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"media_action/api/defs"
	"encoding/json"
	"media_action/api/dbops"
	"media_action/api/session"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//io.WriteString(w, "Create User Handler")
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}
	if err := json.Unmarshal(res, &ubody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	//增加用户 插入到mongo中
	if err := dbops.AddUserCredential(ubody.UserName, ubody.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	//生成session id 写到缓存并且写入到mongo中
	id := session.GenerateNewSessionId(ubody.UserName)
	su := &defs.SignedUp{Success: true, SessionId: id}

	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFail)
		return
	} else {
		sendNormalResponse(w, string(resp), 201)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userName := p.ByName("user_name")
	io.WriteString(w, userName)
}
