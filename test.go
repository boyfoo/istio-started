package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"istio-started/gsrc/pbfiles"
	"log"
)

func main() {
	client, err := grpc.DialContext(context.Background(), ":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	rep := &pbfiles.ProdResponse{}
	err = client.Invoke(context.Background(), "ProdService/GetPord",
		&pbfiles.ProdRequest{
			ProdId: 11123,
		},
		rep,
	)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(rep.Result)
}
