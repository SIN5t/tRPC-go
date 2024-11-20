package post

import (
	"github.com/SIN5t/tRPC-go/app/community/entity"
	"strconv"
)

type Post struct {
	Id       int64  `db:id`
	TopicId  int64  `db:topic_id`
	CreateAt string `db:create_at`
	content  string `db:content`
}

func (p Post) TableName() string {
	return "posts"
}

func (p Post) ToEntity() *entity.Post {
	return &entity.Post{
		Id: strconv.FormatInt(p.Id, 10),
	}
}
