package main

import (
	"github.com/SIN5t/tRPC-go/app/user/service"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
)

func main() {
	s := trpc.NewServer()
	if err := service.RegisterUserService(s); err != nil {
		log.Fatal("server register err %v", err)
	}
	if err := s.Serve(); err != nil {
		log.Fatal("server Serve err %v", err)
	}

}
