package handler

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "Create User Handler")
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}
