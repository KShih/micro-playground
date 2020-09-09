package main

import (
	"fmt"
	"micro-playground/handler"
	proto "micro-playground/proto/helloworld"

	"github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService(
		micro.Name("Greeter"),
	)

	service.Init()

	proto.RegisterGreeterHandler(service.Server(), new(handler.GreeterServiceHandler))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
