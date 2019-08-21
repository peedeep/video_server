package handler

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"io/ioutil"
	"video-server/api/defs"
	"encoding/json"
	"video-server/api/dbops"
	"video-server/api/session"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}

	if err := json.Unmarshal(res, ubody); err != nil {
		SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	err := dbops.AddUserCredential(ubody.Username, ubody.Pwd)
	if err != nil {
		SendErrorResponse(w, defs.ErrorDBError)
		return
	}
	sessionId := session.GenerateNewSessionId(ubody.Username)
	su := &defs.SignedUp{Success: true, SessionId: sessionId}
	if resp, err := json.Marshal(su); err != nil {
		SendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		SendNormalResponse(w, string(resp), 201)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}
	if err := json.Unmarshal(res, &ubody); err != nil {
		SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	pwd, err := dbops.GetUserCredentail(ubody.Username)
	if err != nil {
		SendErrorResponse(w, defs.ErrorDBError)
		return
	}
	if len(pwd) == 0 || ubody.Pwd != pwd {
		SendErrorResponse(w, defs.ErrorNotAuthUser)
	} else {
		sessionId := session.GenerateNewSessionId(ubody.Username)
		login := &defs.Login{Success: true, SessionId: sessionId}
		resp, err := json.Marshal(login)
		if err != nil {
			SendErrorResponse(w, defs.ErrorInternalFaults)
		} else {
			SendNormalResponse(w, string(resp), 201)
		}
	}
}
