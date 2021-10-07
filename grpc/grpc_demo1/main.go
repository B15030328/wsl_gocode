package main

import (
	"grpc/grpc_demo1/services"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	creds, err := credentials.NewServerTLSFromFile("keys/test.crt", "keys/test.pub.key")
	if err != nil {
		log.Fatal(err)
	}
	rpcService := grpc.NewServer(grpc.Creds(creds))
	services.RegisterProductServiceServer(rpcService, new(services.ProducService))

	lis, _ := net.Listen("tcp", ":8091")

	rpcService.Serve(lis)

}
