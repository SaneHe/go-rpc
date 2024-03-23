package server

import (
	"context"
	"sane.com/demo/sane.com/hello"
)

type HelloService struct {
	*hello.UnimplementedHelloServiceServer
}

func (s *HelloService) Hello(_ context.Context, in *hello.HelloRe) (*hello.HelloRe, error) {
	return &hello.HelloRe{
		Name: "hello:" + in.String(),
	}, nil
}
