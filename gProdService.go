package main

import (
	"google.golang.org/grpc"
	"istio-started/gsrc/pbfiles"
	"istio-started/gsrc/service"
	"log"
	"net"
)

func main() {
	myserver := grpc.NewServer()
	pbfiles.RegisterProdServiceServer(myserver, service.NewProdService())
	lis, _ := net.Listen("tcp", ":8080")
	if err := myserver.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
