package handler

import (
	"context"

	proto "next-hello/proto/nextHelloWorld"
)

type NextGreeterServiceHandler struct{}

func (g *NextGreeterServiceHandler) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = " 你好你好, " + req.Name
	return nil
}
