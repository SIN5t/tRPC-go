package topic

import (
	"context"
	"github.com/SIN5t/tRPC-go/proto/community"
	"trpc.group/trpc-go/trpc-database/mysql"
)

type CommunityTopicRepo struct {
	d Dependency
}
type Dependency struct {
	DBGetter func(ctx context.Context) (mysql.Client, error)
}

func (r CommunityTopicRepo) InitRepoTopic(d Dependency) error {
	r.d = d
	return nil
}

func (t *CommunityTopicRepo) QueryTopicById(ctx context.Context, topicId int64) (*community.Topic, error) {
	// 业务逻辑

	return nil, nil
}
