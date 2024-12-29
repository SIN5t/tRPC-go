package topic

import (
	"context"
	"github.com/SIN5t/tRPC-go/app/community/repo"
	"github.com/SIN5t/tRPC-go/proto/community"
	"trpc.group/trpc-go/trpc-database/mysql"
)

type CommunityTopicRepo struct {
	d Dependency
}
type Dependency struct {
	DBGetter func(ctx context.Context) (mysql.Client, error)
}

func (r *CommunityTopicRepo) InitRepoTopic(d Dependency) error {
	r.d = d
	return nil
}

//func (r *CommunityTopicRepo) QueryTopicById(ctx context.Context, topicId int64) (*community.Topic, error) {
//	// 业务逻辑
//
//	db, err := r.d.DBGetter(ctx)
//	if err != nil {
//		fmt.Errorf("db getter error (%w)", err)
//	}
//
//	topicDao := &community.Topic{}
//
//	db.Select(ctx, topicDao, "select * from Topic where topic_id = ?", topicId)
//	return topicDao, nil
//}

func (r *CommunityTopicRepo) QueryTopicById(ctx context.Context, topicId int64) (*community.Topic, error) {
	// 业务逻辑
	db := repo.DB

	topicDao := &Topic{
		ID: 1,
	}
	db.Model(&Topic{}).Find(topicDao)

	topicBui := &community.Topic{}
	return topicBui, nil
}
