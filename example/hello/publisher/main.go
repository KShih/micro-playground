package main

import (
	"context"
	"fmt"
	"time"

	"github.com/micro/go-micro/v2"

	hello "hello/proto/hello"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.foo.srv.hello.pub"), // name the client service
	)
	// Initialise service
	service.Init()

	// create publisher
	pub := micro.NewPublisher("com.foo.service.hello", service.Client())

	// publish message every second
	for now := range time.Tick(time.Second) {
		fmt.Println("now: ", now)
		if err := pub.Publish(context.TODO(), &hello.Message{Say: now.String()}); err != nil {
			// log.Fatal("publish err", err)
		}
	}
}
