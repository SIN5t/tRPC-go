package service

import (
	"context"
	"fmt"
	"strconv"
)
import "github.com/SIN5t/tRPC-go/proto/community"

type communityImpl struct {
	d Dependency
}

type Dependency interface {
	/// repo中有该接口的实现类，那么创建communityImpl就可以用repo的结构体作为成员
	QueryTopicById(ctx context.Context, topicId int64) (*community.Topic, error)
	QueryPostByTopicId(ctx context.Context, req *community.GetPostRequest) (*community.GetPostResponse, error)
}

func (comm *communityImpl) GetTopicById(ctx context.Context, request *community.GetTopicRequest) (*community.GetTopicResponse, error) {

	topicId, err2 := strconv.ParseInt(request.GetId(), 64, 10)
	if err2 != nil {
		fmt.Errorf("转换出错(%w)", err2)
	}
	topic, err := comm.d.QueryTopicById(ctx, topicId)
	commResp := &community.GetTopicResponse{
		ErrCode: 0,
		ErrMsg:  "",
		Topic:   topic,
	}
	return commResp, err
}
