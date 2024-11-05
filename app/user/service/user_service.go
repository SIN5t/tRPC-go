package service

import (
	"context"
	"github.com/SIN5t/tRPC-go/proto/user"
	"trpc.group/trpc-go/trpc-go/server"
)

// 注册服务
type userImpl struct {
}

func NewUserImpl() *userImpl {
	return &userImpl{}
}

func RegisterUserService(s server.Service) error {
	user.RegisterUserService(s, NewUserImpl())
	return nil
}

func (userImpl) GetAccountByUserName(ctx context.Context, req *user.GetAccountByUserNameRequest) (*user.GetAccountByUserNameResponse, error) {

	userMap := map[string]string{"yicwu": "12345", "yicc": "111"}
	pwd := userMap[req.GetUsername()]
	accountByUserNameResponse := &user.GetAccountByUserNameResponse{
		ErrCode:      0,
		ErrMsg:       "",
		UserId:       "",
		Username:     req.GetUsername(),
		PasswordHash: pwd,
		CreateTsSec:  0,
	}
	return accountByUserNameResponse, nil
}
