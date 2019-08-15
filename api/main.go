package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video-server/api/handler"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", handler.CreateUserHandler)
	router.POST("/user/:user_name", handler.Login)
	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8000", r)
}
