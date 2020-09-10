package main

import (
	"context"
	"fmt"
	proto "micro-playground/proto/helloworld"

	"github.com/micro/go-micro/v2"
)

func main() {

	service := micro.NewService(micro.Name("Greeter.Client"))

	service.Init()

	// 建 Greeter 客户端
	greeter := proto.NewGreeterService("Greeter", service.Client())

	// 遠程調用 Greeter 服務的 Hello 方法
	rsp1, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "Jeff"})
	if err != nil {
		fmt.Println(err)
	}

	// 調用 recursive
	rsp2, err := greeter.NextHello(context.TODO(), &proto.HelloRequest{Name: "Alex"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp1.Greeting)
	fmt.Println(rsp2.Greeting)
}
