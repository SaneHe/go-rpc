package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"sane.com/demo/sane.com/hello"
	"time"
)

func Call(n string) {

	// 发起连接，grpc.WithTransportCredentials(insecure.NewCredentials()) 表示使用不安全的连接，即不使用SSL
	conn, err := grpc.Dial("127.0.0.1:9901", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("connect failed: %v", err)
	}

	defer conn.Close()

	c := hello.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	r, err := c.Hello(ctx, &hello.HelloRe{Name: n})
	if err != nil {
		log.Fatalf("call service failed: %v", err)
	}
	println("call service success: ", r.GetName())
}
