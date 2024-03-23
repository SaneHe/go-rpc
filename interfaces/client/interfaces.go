package main

import "net/rpc"

const HelleServiceName = "sane.com/demo/HelloService"

type HelloServiceInterface interface {
	Hello(string, *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelleServiceName, svc)
}
