package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func registerHandler() *httprouter.Router {
	r := httprouter.New()

	r.GET("/videos/:vid-id", streamHandler)
	r.POST("/videos/:vid-id", uploadHandler)

	return r
}
func main() {
	router := registerHandler()
	err := http.ListenAndServe(":8093", router)
	if err != nil {
		log.Println("server begin err:", err)
	}
}
