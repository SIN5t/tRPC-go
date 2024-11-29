package post

import (
	"context"
	"fmt"
	"github.com/SIN5t/tRPC-go/proto/community"
	"trpc.group/trpc-go/trpc-database/mysql"
)

type CommunityPostRepo struct {
	d Dependency
}

type Dependency struct {
	DBGetter func(ctx context.Context) (mysql.Client, error)
}

func (p *CommunityPostRepo) InitPostRepo(dependency Dependency) error {
	p.d = dependency
	return nil
}

//QueryPostByTopicId(ctx context.Context, topicId int64) (*community.Post, error)

func (p *CommunityPostRepo) QueryPostByTopicId(ctx context.Context, topicId int64) (*community.Post, error) {

	// db逻辑
	db, err := p.d.DBGetter(ctx)
	if err != nil {
		fmt.Errorf("d.DBGetter() err:%s", err)
		return nil, err
	}
	postDao := &community.Post{}
	db.Select(ctx, postDao, "select * from post where topic_id = ?", topicId)
	return postDao, nil

}
