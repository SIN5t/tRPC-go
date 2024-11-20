package repo

import (
	"context"
	"github.com/SIN5t/tRPC-go/app/community/repo/topic"
	"trpc.group/trpc-go/trpc-database/mysql"
)

type Repo struct {
	t topic.CommunityTopicRepo
}

type Dependency struct {
	ClientName string
}

func NewRepo(d Dependency) (*Repo, error) {
	var dbGetter = func(ctx context.Context) (mysql.Client, error) {
		return mysql.NewUnsafeClient(d.ClientName), nil
	}

	r := &Repo{}
	if err := r.t.InitRepoTopic(topic.Dependency{DBGetter: dbGetter}); err != nil {
		return nil, err
	}

	return r, nil
}
