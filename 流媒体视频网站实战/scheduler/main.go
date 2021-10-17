package main

import (
	"log"
	"net/http"
	"stream/scheduler/dbops"
	"stream/scheduler/taskrunner"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/video-delete-record/:vid-id", vidDelRecHandler)

	return router
}

func main() {
	dbops.Init()
	go taskrunner.Start()
	log.Println("sheduler server start")
	r := RegisterHandlers()
	http.ListenAndServe(":9001", r)
}
