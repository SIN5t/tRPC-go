package post

import (
	"context"
	"github.com/SIN5t/tRPC-go/app/community/repo"
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

//// QueryPostByTopicId (ctx context.Context, topicId int64) (*community.Post, error)
//func (p *CommunityPostRepo) QueryPostByTopicId(ctx context.Context, topicId int64) (*community.Post, error) {
//
//	// db逻辑
//	db, err := p.d.DBGetter(ctx)
//	if err != nil {
//		fmt.Errorf("d.DBGetter() err:%s", err)
//		return nil, err
//	}
//	postDao := &community.Post{}
//	postDao1 := &Post{} // db相关的
//	db.Select(ctx, postDao1, "select * from post where topic_id = ?", topicId)
//
//	// todo dao和业务定义的结构体映射一下过去
//	return postDao, nil
//
//}

func (p *CommunityPostRepo) QueryPostByTopicId(ctx context.Context, topicId int64) (*community.Post, error) {

	// db逻辑
	db := repo.DB
	postDao := &Post{
		TopicId: topicId,
	} // db相关的
	db.Model(&Post{}).Find(postDao)
	postLogic := &community.Post{}

	// todo dao和业务定义的结构体映射一下过去
	return postLogic, nil

}
