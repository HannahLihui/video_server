package session

import (
	"github.com/HannahLihui/video_server/video_server/api/utils"
	"sync"
	"time"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}
func LoadSessionsFromDB() {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}
	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}
func GenerateNewSessionId(un string) string {
	id, _ := utils.NewUUID()
	ct := time.Now().UnixNano() / 1000000
	ttl := ct * 30 * 1000
	ss := &defs.SimpleSession{
		UserName: un,
		TTL:      ttl,
	}
	sessionMap.Store(id, ss)
	dbops.InsertSession(id, ttl, un)
	return id
}
func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(id)
	if ok {
		st := time.Now().UnixNano() / 1000000
		if ss(*defs.SimpleSession{}).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).UserName, false
	}
	return "", true

}
func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeletedSession(sid)
}
