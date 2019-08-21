package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video-server/api/handler"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//check session
	isExpired := handler.ValidateUserSession(r)
	if isExpired {
		//重新登录
	} else {
		//已经登录
	}
	m.r.ServeHTTP(w, r)
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", handler.CreateUserHandler)
	router.POST("/user/:user_name", handler.Login)
	return router
}

func main() {
	r := RegisterHandlers()
	wareHandler := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", wareHandler)
}
