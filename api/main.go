package main

import (
	"context"
	"encoding/json"
	"log"
	hello "micro-playground/proto/helloworld"
	"strings"

	grmon "github.com/bcicen/grmon/agent"
	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/v2"
)

type Say struct {
	Client hello.GreeterService
}

// Hello => Will listen on /say/hello/ or /Say/Hello/
func (s *Say) Hello(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("收到 Say.Hello API 請求") // Full request e.g.: http://localhost:8080/greeter/Say/Hello?name=jeff

	// 從request.GET 中獲取name值
	name, ok := req.Get["name"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.greeter", "名字不能為空")
	}

	// 將餐數交給底層服務處理
	response, err := s.Client.Hello(ctx, &hello.HelloRequest{
		Name: strings.Join(name.Values, " "),
	})
	if err != nil {
		return err
	}

	// 處理成功則返回結果
	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": response.Greeting,
	})
	rsp.Body = string(b)

	return nil
}

func (s *Say) NextHello(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("收到 Say.Hello API 請求")

	name, ok := req.Get["name"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.greeter", "名字不能為空")
	}

	response, err := s.Client.NextHello(ctx, &hello.HelloRequest{
		Name: strings.Join(name.Values, " "),
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": response.Greeting,
	})
	rsp.Body = string(b)

	return nil
}

func main() {
	grmon.Start()
	service := micro.NewService(
		micro.Name("go.micro.api.greeter"), // visit => http://localhost:8080/greeter
	)

	service.Init()

	// 將參數轉給底層service Greeter 處理
	service.Server().Handle(
		service.Server().NewHandler(
			&Say{Client: hello.NewGreeterService("Greeter", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
