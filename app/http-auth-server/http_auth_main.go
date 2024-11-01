package main

import (
	"github.com/SIN5t/tRPC-go/app/http-auth-server/service"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
)

func main() {
	s := trpc.NewServer()
	if err := service.RegisterHttpAuthService(s); err != nil {
		log.Fatal("服务初始化失败: %v", err)
	}
	if err := s.Serve(); err != nil {
		log.Fatal("服务开启失败： %v", err)
	}

}
