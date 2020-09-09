package subscriber

// 實現異步消息接收並處理的地方。其中展示了用兩種不同方式處理消息，一是以對象方法處理， 二是以一個函數來處理。

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	hello "hello/proto/hello"
)

type Hello struct{}

func (e *Hello) Handle(ctx context.Context, msg *hello.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *hello.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
