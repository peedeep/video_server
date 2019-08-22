package handler

import (
	"net/http"
	"io"
)

func SendErrorResponse(w http.ResponseWriter, SC int, err string) {
	w.WriteHeader(SC)
	io.WriteString(w, err)
}
