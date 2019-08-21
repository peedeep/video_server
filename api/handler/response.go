package handler

import (
	"net/http"
	"video-server/api/defs"
	"encoding/json"
	"io"
)

func SendErrorResponse(w http.ResponseWriter, errResp defs.ErrorResponse) {
	w.WriteHeader(errResp.HttpSC)
	bytes, _ := json.Marshal(&errResp.Error)
	io.WriteString(w, string(bytes))

}

func SendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
