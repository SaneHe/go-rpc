package main

import (
	"net"
	"net/rpc"
)

func main() {

	_ = RegisterHelloService(new(HelloService))
	listener, err := net.Listen("tcp", "127.0.0.1:9901")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go rpc.ServeConn(conn)
	}
}
