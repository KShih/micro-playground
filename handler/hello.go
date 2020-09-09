package handler

import (
	"context"

	proto "micro-playground/proto/helloworld"
)

type GreeterServiceHandler struct{}

func (g *GreeterServiceHandler) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = " 你好, " + req.Name
	return nil
}
