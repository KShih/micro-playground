package handler

import (
	"context"
	"fmt"

	nextProto "micro-playground/proto/nextHelloWorld"

	proto "micro-playground/proto/helloworld"
)

type GreeterServiceHandler struct {
	NextHelloClient nextProto.NextGreeterService
}

func (g *GreeterServiceHandler) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	var val string
	x := req

	if x, ok := x.GetValue().(*proto.HelloRequest_Texter); ok {
		val = x.Texter
	}

	if x, ok := x.GetValue().(*proto.HelloRequest_Number); ok {
		val = fmt.Sprintf("%f", x.Number)
	}

	rsp.Greeting = " 你好, " + req.Name + val

	// for _, tag := range req.GetTagList() {
	// 	rsp.TagResp += tag.GetTableName() + "\n"
	// 	rsp.TagResp += tag.GetTableName() + "\n"
	// 	rsp.TagResp += tag.GetTagName() + "\n"
	// 	rsp.TagResp += tag.GetTagType() + "\n"
	// 	rsp.TagResp += "-------\n"
	// }

	info2 := &proto.Info2{}
	//info2.ConnectType = &proto.Info2_ConnectString{ConnectString: "resp connectString"}
	info2.ConnectType = &proto.Info2_DsnNmae{DsnNmae: "DSN_NAME"}
	//info2.TagList = req.TagList
	info2.DbType = "INFO2_DB_TYPE"

	rsp.Info = &proto.HelloResponse_Info2{Info2: info2}
	fmt.Println(rsp.Greeting)

	return nil
}

func (g *GreeterServiceHandler) NextHello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {

	nextRsp, err := g.NextHelloClient.Hello(ctx, &nextProto.HelloRequest{Name: req.Name}) // take the input and pass to another service
	rsp.Greeting = nextRsp.Greeting
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

/*
{
	"name":"jeff",
	"number":2.3,
	"tagList":[
		 {
				"tableName":"tn",
				"tagName":"tg",
				"columnName":"cn",
				"tagType":"tt"
		 },
		 {
				"tableName":"tn",
				"tagName":"tg",
				"columnName":"cn",
				"tagType":"tt"
		 },
		 {
				"tableName":"tn",
				"tagName":"tg",
				"columnName":"cn",
				"tagType":"tt"
		 }
	]
}
*/
