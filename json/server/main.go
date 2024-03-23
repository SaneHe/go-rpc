package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// {"method":"HelloService.Hello","params":["sane"],"id":0}
func main() {
	_ = rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", "127.0.0.1:9901")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
