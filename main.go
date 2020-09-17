package main

import (
	"fmt"
	"micro-playground/handler"
	proto "micro-playground/proto/helloworld"
	nextProto "micro-playground/proto/nextHelloWorld"

	grmon "github.com/bcicen/grmon/agent"
	"github.com/micro/go-micro/v2"
)

func main() {
	grmon.Start()
	service := micro.NewService(
		micro.Name("Greeter"),
	)

	service.Init()

	proto.RegisterGreeterHandler(service.Server(), &handler.GreeterServiceHandler{
		// call service-B in service-A by client-A
		// init the handler's Client by instantiating the exist service named "NextGreeter"
		NextHelloClient: nextProto.NewNextGreeterService("NextGreeter", service.Client()),
	})

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
