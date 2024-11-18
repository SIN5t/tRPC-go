package main

import (
	"github.com/SIN5t/tRPC-go/app/user/repo"
	"github.com/SIN5t/tRPC-go/app/user/service"
	"trpc.group/trpc-go/trpc-go"
)

func main() {
	s := trpc.NewServer()
	r, err := initRepo()
	if err != nil {
		return
	}
	if err := service.RegisterUserService(s, r); err != nil {
		return
	}
	s.Serve()

}

func initRepo() (*repo.Repo, error) {
	d := repo.Dependency{UserAccountDBClientName: "db."}
	r, err := repo.NewRepo(d)
	if err != nil {
		return nil, err
	}
	return r, nil
}
