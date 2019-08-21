package handler

import (
	"net/http"
	"video-server/api/session"
	"video-server/api/defs"
)

var HeaderFieldSession = "X-Session-Id"
var HeaderFieldUsername = "X-User-Name"

func ValidateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HeaderFieldSession)
	if len(sid) == 0 {
		return false
	}
	uname, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}
	r.Header.Add(HeaderFieldUsername, uname)
	return true
}

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HeaderFieldUsername)
	if len(uname) == 0 {
		SendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}
	return true
}