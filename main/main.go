package main

import (
	"context"
	simplest "tRPC-go/proto"
	"time"
	"trpc.group/trpc-go/trpc-go"
)

func main() {
	server := trpc.NewServer()
	simplest.RegisterHelloWorldService(server, helloWorldImpl{})
	server.Serve()
}

type helloWorldImpl struct {
}

func (h helloWorldImpl) Hello(ctx context.Context, req *simplest.HelloRequest) (*simplest.HelloResponse, error) {
	greetingReq := req.Greeting

	resp := &simplest.HelloResponse{
		ErrCode:   200,
		ErrMsg:    "",
		Response:  "",
		Timestamp: 0,
	}
	resp.Response = "greeting: " + greetingReq + "hello !!"
	resp.Timestamp = float64(time.Now().UnixMilli())
	return resp, nil
}
