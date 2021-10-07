package main

import (
	"context"
	"fmt"
	"grpc/grpc_demo1/services"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("../keys/test.crt", "test.com")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(":8091", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	producClient := services.NewProductServiceClient(conn)
	resp, err := producClient.GetProductStock(context.Background(), &services.ProdRequest{ProdId: 2})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.ProdStock)

}
