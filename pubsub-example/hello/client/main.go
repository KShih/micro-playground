package main

import (
	"context"
	"fmt"
	hello "hello/proto/hello"

	"github.com/micro/go-micro/v2"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.foo.service.hello.client"), //name the client service
	)
	// Initialise service
	service.Init()

	//create hello service client
	helloClient := hello.NewHelloService("com.foo.service.hello", service.Client())

	//invoke hello service method
	resp, err := helloClient.Call(context.TODO(), &hello.Request{Name: "Bill 4"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.Msg)
}
