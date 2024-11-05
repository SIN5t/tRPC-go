package service

import (
	"context"
	"fmt"
	"github.com/SIN5t/tRPC-go/proto/common"
	"github.com/SIN5t/tRPC-go/proto/http_auth"
	"github.com/SIN5t/tRPC-go/proto/user"
	"trpc.group/trpc-go/trpc-go/server"
)

// 服务注册
type httpAuthImpl struct {
}

func newHttpAuthImpl() httpAuthImpl {
	return httpAuthImpl{}
}

func RegisterHttpAuthService(s server.Service) error {
	httpImpl := newHttpAuthImpl()
	http_auth.RegisterAuthService(s, httpImpl)
	return nil
}

func (httpAuthImpl) Login(ctx context.Context, req *http_auth.LoginRequest) (*http_auth.LoginResponse, error) {
	context.WithValue(ctx, "loginTest", "value")
	username, pwd := req.GetUsername(), req.GetPasswordHash()
	traceId := req.GetMataData().GetTraceId()
	fmt.Println(username)
	fmt.Println(pwd)

	userNameRequest := &user.GetAccountByUserNameRequest{
		Username: username,
		MetaData: &common.MetaData{TraceId: username},
	}
	userNameResponse, err := user.NewUserClientProxy().GetAccountByUserName(ctx, userNameRequest)
	loginResponse := &http_auth.LoginResponse{
		ErrCode: 0,
		ErrMsg:  "",
		Data:    &http_auth.LoginResponse_Data{IdTicket: traceId},
	}

	// 业务逻辑，判断密码是否正确

	if userNameResponse.PasswordHash != req.GetPasswordHash() {
		loginResponse.ErrCode = -1
		loginResponse.ErrMsg = "密码错误"
	}

	return loginResponse, err
}
