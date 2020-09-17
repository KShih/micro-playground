package handler

import (
	"context"
	"fmt"

	nextProto "micro-playground/proto/nextHelloWorld"

	proto "micro-playground/proto/helloworld"
)

type GreeterServiceHandler struct {
	NextHelloClient nextProto.NextGreeterService
}

func (g *GreeterServiceHandler) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = " 你好, " + req.Name
	return nil
}

func (g *GreeterServiceHandler) NextHello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {

	nextRsp, err := g.NextHelloClient.Hello(ctx, &nextProto.HelloRequest{Name: req.Name}) // take the input and pass to another service
	rsp.Greeting = nextRsp.Greeting
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
