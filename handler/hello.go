package handler

import (
	"context"
	"fmt"

	nextProto "micro-playground/proto/nextHelloWorld"

	proto "micro-playground/proto/helloworld"

	"github.com/micro/go-micro/v2"
)

type GreeterServiceHandler struct{}

func (g *GreeterServiceHandler) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = " 你好, " + req.Name
	return nil
}

func (g *GreeterServiceHandler) NextHello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	service := micro.NewService(micro.Name("Greeter.Client"))

	service.Init()

	// finding another service
	nextGreeter := nextProto.NewNextGreeterService("NextGreeter", service.Client())
	nextRsp, err := nextGreeter.Hello(context.TODO(), &nextProto.HelloRequest{Name: req.Name}) // take the input and pass to another service
	rsp.Greeting = nextRsp.Greeting
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Greeting)

	// rsp.Greeting =
	return nil
}
