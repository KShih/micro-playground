package main

import (
	"fmt"
	"next-hello/handler"
	proto "next-hello/proto/nextHelloWorld"

	"github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService(
		micro.Name("NextGreeter"),
	)

	service.Init()

	proto.RegisterNextGreeterHandler(service.Server(), new(handler.NextGreeterServiceHandler))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
