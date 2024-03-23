package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"sane.com/demo/client"
	"sane.com/demo/sane.com/hello"
	"sane.com/demo/server"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:9901")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	hello.RegisterHelloServiceServer(s, &server.HelloService{})

	// reflection.Register(s)
	fmt.Println("gRPC server listen in 9901...")
	go func() {
		client.Call("sane")
	}()
	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
