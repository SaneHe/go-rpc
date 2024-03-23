package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// curl localhost:9901/json-rpc -X POST --data '{"method":"HelloService.Hello","params":["hello"],"id":0}'
// {"id":0,"result":"hello:hello","error":null}

func main() {
	_ = rpc.RegisterName("HelloService", new(HelloService))

	http.HandleFunc("/json-rpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}

		err := rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	err := http.ListenAndServe("127.0.0.1:9901", nil)
	if err != nil {
		panic(err)
	}
}
