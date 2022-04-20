package main

import (
	"fmt"
	"gokit1/consul"
	"gokit1/gokitService"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	kithttp "github.com/go-kit/kit/transport/http"
	mymux "github.com/gorilla/mux"
)

func main() {
	var service gokitService.UserService
	endp := gokitService.GenUserEndpoint(service)
	UsernameServer := kithttp.NewServer(endp, gokitService.DecodeUserNameRequest, gokitService.EncodeUserNameResponse)

	r := mymux.NewRouter()
	r.Methods("GET", "DELETE").Path(`/user/{uid:\d+}`).Handler(UsernameServer)
	r.Methods("GET").Path("/health").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})

	errChan := make(chan error)

	// 注册服务
	go func() {
		consul.RegService() //调用注册服务程序
		err := http.ListenAndServe(":8092", r)
		if err != nil {
			errChan <- err
		}
	}()

	// 注销服务
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM) //监听到interrupt或kill时写进channel
		errChan <- fmt.Errorf("%s", <-sigChan)
	}()

	geterr := <-errChan
	consul.DeRegService()
	log.Println(geterr)
}
