package session

import (
	"sync"
	"media_action/api/dbops"
	"media_action/api/defs"
	"media_action/api/utils"
)

var sessionMap *sync.Map //线程安全 ，读很快，写有问题要加锁

func init() {
	sessionMap = &sync.Map{}

}

//每次重启从db里同步到cache中
func LoadSessionsFromDB() {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}
	//将返回来的数据存储到全局的sync.map去
	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}

// 生成sessionid
func GenerateNewSessionId(userName string) string {
	id, _ := utils.NewUUID()
	ct := defs.GetMillTime()
	ttl := ct + 30*60*1000 // 过期时间设置为30分钟

	ss := &defs.SimpleSession{SessionId: id, UserName: userName, TTL: ttl}
	sessionMap.Store(id, ss)

	dbops.InsertSession(id, ttl, userName)
	return id
}

// 是否过期
// username ，err
func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := defs.GetMillTime()
		if ss.(*defs.SimpleSession).TTL < ct {
			//delete  expired session  (cache and db)
			DeleteExpiredSession(sid)
			return "", true //过期
		}
		return ss.(*defs.SimpleSession).UserName, false
	}
	return "", true
}

func DeleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}
