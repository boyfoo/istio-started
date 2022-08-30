package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"istio-started/gsrc/pbfiles"
	"log"
)

func main() {
	client, err := grpc.DialContext(context.Background(), "grpc.jtthink.com:31259", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}
	rsp := &pbfiles.ProdResponse{}
	err = client.Invoke(context.Background(),
		"/ProdService/GetProd",
		&pbfiles.ProdRequest{ProdId: 123}, rsp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rsp.Result)
}
