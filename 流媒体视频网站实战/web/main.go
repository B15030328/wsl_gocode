package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func registerHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", homeHandler)
	router.POST("/", homeHandler)
	router.GET("/userhome", userHomeHandler)
	router.POST("/userhome", userHomeHandler)
	router.POST("/api", apiHandler)
	router.POST("/upload/:vid-id", proxyHandler)
	router.ServeFiles("/statics/*filepath", http.Dir("../template"))
	return router
}

func main() {
	r := registerHandler()
	log.Println("web server begin success,port:8080")
	http.ListenAndServe(":8080", r)
}
