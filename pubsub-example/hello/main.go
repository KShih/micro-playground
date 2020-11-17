package main

import (
	"hello/handler"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/server"

	hello "hello/proto/hello"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.foo.service.hello"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	hello.RegisterHelloHandler(service.Server(), new(handler.Hello)) // 一個服務中可以註冊多個Handler以完成不同業務功能

	// Register Struct as Subscriber 消息處理
	// 群播
	//micro.RegisterSubscriber("com.foo.service.hello", service.Server(), new(subscriber.Hello))

	// 從queue取,
	micro.RegisterSubscriber("com.foo.service.hello", service.Server(), new(handler.SubHello), server.SubscriberQueue("tester-queue"))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
