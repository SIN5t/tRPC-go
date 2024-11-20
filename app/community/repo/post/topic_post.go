package post

import (
	"context"
	"fmt"
	"github.com/SIN5t/tRPC-go/proto/community"
	"trpc.group/trpc-go/trpc-database/mysql"
)

type PostRepo struct {
	d Dependency
}

type Dependency struct {
	DBGetter func(ctx context.Context) (mysql.Client, error)
}

func (p PostRepo) InitPostRepo(dependency Dependency) {
	p.d = dependency
}

//QueryPostByTopicId(ctx context.Context, topicId int64) (*community.Post, error)

func (p *PostRepo) QueryPostByTopicId(ctx context.Context, topicId int64) (*community.Post, error) {

	// db逻辑
	db, err := p.d.DBGetter(ctx)
	if err != nil {
		fmt.Errorf("d.DBGetter() err:%s", err)
		return nil, err
	}
	postDao := &community.Post{}
	db.Select(ctx, postDao, "selct * from post where topic_id = ?", topicId)
	return postDao, nil

}
