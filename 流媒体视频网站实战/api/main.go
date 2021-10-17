package main

import (
	"net/http"
	"stream/api/dbopts"
	"stream/api/session"

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
	r.GET("/user/:username", GetUserInfo)
	r.POST("/user/:username/videos", AddNewVideo)
	r.GET("/user/:username/videos", ListAllVideos)
	r.DELETE("/user/:username/videos/:vid-id", DeleteVideo)
	r.POST("/videos/:vid-id/comments", PostComment)
	r.GET("/videos/:vid-id/comments", ShowComments)
	return r

}

// func Prepare() {
// 	session.LoadSessionsFromDB()
// }
func main() {
	dbopts.Init()
	session.Init()
	router := registerRouter()
	mh := NewMiddlewareHandler(router)
	http.ListenAndServe(":8000", mh)
}
