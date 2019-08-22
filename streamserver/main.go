package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video-server/streamserver/handler"
)

type middleWareHandler struct {
	r *httprouter.Router
	l *handler.ConnLimiter
}

func NewMiddleWareHandler(r *httprouter.Router, l *handler.ConnLimiter) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = l
	return m
}

func registerHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/videos/:video-id", handler.StreamHandler)
	router.POST("/upload/:video-id", handler.UploadHandler)
	router.POST("/testpage", handler.TestPageHandler)
	return router
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() {
		handler.SendErrorResponse(w, http.StatusTooManyRequests, "Too many requests")
		return
	}
	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

func main() {
	r := registerHandlers()
	h := NewMiddleWareHandler(r, handler.NewConnLimiter(2))
	http.ListenAndServe(":9000", h)
}
