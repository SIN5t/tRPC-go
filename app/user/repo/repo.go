package repo

import (
	"context"
	"fmt"
	"github.com/SIN5t/tRPC-go/app/user/repo/account"
	"trpc.group/trpc-go/trpc-database/mysql"
)

type Repo struct {
	account.UserAccountRepository
}
type Dependency struct {
	UserAccountDBClientName string
}

// NewRepo 新建 user 服务所需的 repo 依赖全集
func NewRepo(d Dependency) (*Repo, error) {

	repo := &Repo{}

	// 初始化用户仓库
	accountDep := account.Dependency{
		DBGetter: func(ctx context.Context) (mysql.Client, error) {
			return mysql.NewUnsafeClient(d.UserAccountDBClientName), nil
		},
	}
	if err := repo.InitUserAccountRepository(accountDep); err != nil {
		return nil, fmt.Errorf("初始化用户仓库失败(%w)", err)
	}

	return repo, nil

}
