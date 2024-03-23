package main

import "net/rpc"

func main() {
	conn, err := rpc.Dial("tcp", "127.0.0.1:9901")
	if err != nil {
		panic(err)
	}

	var reply string
	err = conn.Call("HelloService.Hello", "sane", &reply)
	if err != nil {
		panic(err)
	}

	println(reply)
}
