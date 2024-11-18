package account

import (
	"github.com/SIN5t/tRPC-go/app/user/entity"
	"strconv"
)

type userAccountItem struct {
	Id           int64  `db:"id"`
	Username     string `db:username`
	PasswordHash string `db:password_hash` // go中是驼峰，数据库中是蛇形
}

func (userAccountItem) TableName() string {
	return "t_user_account"
}

// ToEntity 将数据库表转为实际
func (u userAccountItem) ToEntity() *entity.Account {
	return &entity.Account{
		ID:       strconv.FormatInt(u.Id, 36),
		Username: u.Username,
		Password: u.PasswordHash,
	}
}
