// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/hello/hello.proto

// 由前文提到的 protoc-gen-micro 生成的， 進一步簡化開發者的工作。其中定義了HelloSerivce 接口， 以及 HelloHandler 接口。後者是我們需要去實現、完成業務邏輯的接口

package com_foo_service_hello

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Hello service

func NewHelloEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Hello service

type HelloService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Hello_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Hello_PingPongService, error)
}

type helloService struct {
	c    client.Client
	name string
}

func NewHelloService(name string, c client.Client) HelloService {
	return &helloService{
		c:    c,
		name: name,
	}
}

func (c *helloService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Hello.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Hello_StreamService, error) {
	req := c.c.NewRequest(c.name, "Hello.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &helloServiceStream{stream}, nil
}

type Hello_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type helloServiceStream struct {
	stream client.Stream
}

func (x *helloServiceStream) Close() error {
	return x.stream.Close()
}

func (x *helloServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *helloServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *helloServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *helloServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloService) PingPong(ctx context.Context, opts ...client.CallOption) (Hello_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Hello.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &helloServicePingPong{stream}, nil
}

type Hello_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type helloServicePingPong struct {
	stream client.Stream
}

func (x *helloServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *helloServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *helloServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *helloServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *helloServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *helloServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Hello service

type HelloHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Hello_StreamStream) error
	PingPong(context.Context, Hello_PingPongStream) error
}

func RegisterHelloHandler(s server.Server, hdlr HelloHandler, opts ...server.HandlerOption) error {
	type hello interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Hello struct {
		hello
	}
	h := &helloHandler{hdlr}
	return s.Handle(s.NewHandler(&Hello{h}, opts...))
}

type helloHandler struct {
	HelloHandler
}

func (h *helloHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.HelloHandler.Call(ctx, in, out)
}

func (h *helloHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.HelloHandler.Stream(ctx, m, &helloStreamStream{stream})
}

type Hello_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type helloStreamStream struct {
	stream server.Stream
}

func (x *helloStreamStream) Close() error {
	return x.stream.Close()
}

func (x *helloStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *helloStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *helloStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *helloStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *helloHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.HelloHandler.PingPong(ctx, &helloPingPongStream{stream})
}

type Hello_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type helloPingPongStream struct {
	stream server.Stream
}

func (x *helloPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *helloPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *helloPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *helloPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *helloPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *helloPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
