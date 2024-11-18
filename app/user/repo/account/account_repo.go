package account

import (
	"context"
	"fmt"
	"github.com/SIN5t/tRPC-go/app/user/entity"
)
import "trpc.group/trpc-go/trpc-database/mysql"

type UserAccountRepository struct {
	dep Dependency
}

type Dependency struct {
	DBGetter func(context.Context) (mysql.Client, error)
}

func (r UserAccountRepository) InitUserAccountRepository(dep Dependency) error {
	r.dep = dep
	return nil
}

// QueryAccountByUsername  业务逻辑，实现service的interface中的方法
func (r UserAccountRepository) QueryAccountByUsername(ctx context.Context, username string) (*entity.Account, error) {
	dbClient, err := r.dep.DBGetter(ctx)
	if err != nil {
		return nil, fmt.Errorf("获取DB失败（%w）", err)
	}
	var userAccountItems []userAccountItem
	query := fmt.Sprintf("select * from %s where username = ? LIMIT 1", userAccountItem{}.TableName())
	if err := dbClient.Select(ctx, &userAccountItems, query, username); err != nil {
		return nil, fmt.Errorf("查询db失败(%w)", err)
	}
	if len(userAccountItems) == 0 {
		return nil, nil
	}
	return userAccountItems[0].ToEntity(), nil
}
