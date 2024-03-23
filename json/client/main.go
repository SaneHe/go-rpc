package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:9901")
	if err != nil {
		panic(err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call("HelloService.Hello", "sane", &reply)
	if err != nil {
		panic(err)
	}

	println(reply)
}
