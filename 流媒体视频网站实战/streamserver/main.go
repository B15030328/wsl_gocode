package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// 流量控制中间件
type middlewareHandler struct {
	r  *httprouter.Router
	cl *ConnLimiter
}

func NewMiddlewareHandler(r *httprouter.Router, cc int) *middlewareHandler {
	mh := &middlewareHandler{
		r:  r,
		cl: NewConnLimiter(cc),
	}
	return mh
}

// 流量控制
func (mh *middlewareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !mh.cl.getConn() {
		sendErrorResponse(w, http.StatusTooManyRequests, "too many Request")
		return
	}
	mh.r.ServeHTTP(w, r)
	defer mh.cl.releaseConn()
}

func registerHandler() *httprouter.Router {
	r := httprouter.New()

	r.GET("/videos/:vid-id", streamHandler)
	r.POST("/videos/:vid-id", uploadHandler)

	return r
}

func main() {
	file2Bytes("xizi.mp4")

	router := registerHandler()
	mh := NewMiddlewareHandler(router, 2)
	log.Println("server begin success, port:8093")
	err := http.ListenAndServe(":8093", mh)
	if err != nil {
		log.Println("server begin err:", err)
	}
}
