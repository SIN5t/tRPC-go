package post

import (
	"github.com/SIN5t/tRPC-go/app/community/entity"
	"strconv"
)

type Post struct {
	ID      int64  `db:id`
	TopicId int64  `db:topic_id`
	Content string `db:Content`
}

func (p Post) TableName() string {
	return "posts"
}

func (p Post) ToEntity() *entity.Post {
	return &entity.Post{
		Id: strconv.FormatInt(p.ID, 10),
	}
}
