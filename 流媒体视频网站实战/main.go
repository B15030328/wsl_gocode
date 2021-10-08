package main

import (
	"net/http"
	"stream/dbopts"
	"stream/session"

	"github.com/julienschmidt/httprouter"
)

type middlewareHandler struct {
	r *httprouter.Router
}

func NewMiddlewareHandler(r *httprouter.Router) http.Handler {

	middleware := &middlewareHandler{}
	middleware.r = r
	return middleware
}

func (m middlewareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//check session
	ValidateUserSession(r)

	m.r.ServeHTTP(w, r)
}

func registerRouter() *httprouter.Router {
	r := httprouter.New()
	r.POST("/user", CreateUser)
	r.POST("/user/:user_name", Login)

	return r

}

func main() {
	dbopts.Init()
	session.Init()
	router := registerRouter()
	mh := NewMiddlewareHandler(router)
	http.ListenAndServe(":8092", mh)
}
