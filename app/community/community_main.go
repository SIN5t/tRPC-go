package main

import (
	"github.com/SIN5t/tRPC-go/app/community/repo"
	"github.com/SIN5t/tRPC-go/app/community/service"
	"trpc.group/trpc-go/trpc-go"
)

func main() {
	s := trpc.NewServer()
	r := newDBGetter()
	err := service.RegisterCommunityService(s, r)
	if err != nil {
		return
	}
	s.Serve()
}

func newDBGetter() *repo.Repo {
	r, err := repo.NewRepo(repo.Dependency{ClientName: "db.mysql.communityTopic"}) // trpc 的连接客户端
	if err != nil {
		return nil
	}
	return r
}
