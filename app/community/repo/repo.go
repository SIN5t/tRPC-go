package repo

import (
	"context"
	"github.com/SIN5t/tRPC-go/app/community/repo/post"
	"github.com/SIN5t/tRPC-go/app/community/repo/topic"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"trpc.group/trpc-go/trpc-database/mysql"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:123456@tcp(127.0.0.1:3307)/yicwu?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(gormMysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	DB = db
}

type Repo struct {
	topic.CommunityTopicRepo
	post.CommunityPostRepo
}

type Dependency struct {
	ClientName string
}

func NewRepo(d Dependency) (*Repo, error) {
	var dbGetter = func(ctx context.Context) (mysql.Client, error) {
		return mysql.NewUnsafeClient(d.ClientName), nil
	}

	r := &Repo{}
	if err := r.InitRepoTopic(topic.Dependency{DBGetter: dbGetter}); err != nil {
		return nil, err
	}
	if err := r.InitPostRepo(post.Dependency{DBGetter: dbGetter}); err != nil {
		return nil, err
	}

	return r, nil
}
