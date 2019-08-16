package session

import (
	"sync"
	"video-server/api/dbops"
	"video-server/api/defs"
	"video-server/api/utils"
	"time"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func NowInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

func DeleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

func LoadSessionsFromDB() {
	sessions, err := dbops.RetriveAllSessions()
	if err != nil {
		return
	}
	sessions.Range(func(key, value interface{}) bool {
		ss := value.(*defs.SimpleSession)
		sessionMap.Store(key, ss)
		return true
	})
	//sessionMap = sessions
}

func GenerateNewSessionId(username string) string {
	sid, err := utils.NewUUID()
	if err != nil {
		return ""
	}
	ttl := NowInMilli() + 30*60*1000
	ss := &defs.SimpleSession{
		Username: username,
		TTL:      ttl,
	}
	sessionMap.Store(sid, ss)
	dbops.InsertSession(sid, ttl, username)
	return sid
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		if NowInMilli() > ss.(*defs.SimpleSession).TTL {
			DeleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).Username, false
	}
	session, err := dbops.RetrieveSession(sid)
	if err != nil || NowInMilli() > session.TTL {
		return "", true
	}
	return session.Username, false
}
