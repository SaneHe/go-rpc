package main

import "net/rpc"

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

type HelloServiceClient struct {
	*rpc.Client
}

func NewHelloServiceClient(network, address string) (*HelloServiceClient, error) {
	client, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}

	return &HelloServiceClient{
		client,
	}, nil
}

func (c *HelloServiceClient) Hello(request string, reply *string) error {
	return c.Client.Call(HelleServiceName+".Hello", request, reply)
}
