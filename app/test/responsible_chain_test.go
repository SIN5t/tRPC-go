package test

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

// 来一个请求，携带对应的权重
// 根据权重，一级一级地走责任链

type Request struct {
	ID     int
	Weight int
	MSG    string
}

// 公共的抽象节点
type Handler interface {
	Handle(request *Request) error
	NextNode(node Handler)
}

// 节点1： mt 节点
type MTHandler struct {
	next Handler
	//suspend bool
}

func (mt *MTHandler) Handle(request *Request) error {
	if request.Weight > 0 && request.Weight < 10 {
		// 向mt发送一个审批，这里可以用channel阻塞mt的审批
		fmt.Printf("mt handler is processing req: %d\n", request.ID)
		// 如果出现异常，整个流程结束
	} else if mt.next != nil {
		mt.next.Handle(request)
	} else {
		fmt.Println("无节点可处理")
		return errors.New("无可处理节点")
	}
	return nil
}

func (mt *MTHandler) NextNode(node Handler) {
	mt.next = node
}

// 节点2： ld 节点
type LDHandler struct {
	next    Handler
	suspend bool
}

func (ld *LDHandler) Handle(request *Request) error {
	if request.Weight >= 10 && request.Weight < 30 {
		// 向mt发送一个审批消息，这里可以用channel阻塞mt的审批
		// 也可以使用数据库，这里卡着隔一段时间扫描数据库，查看审批状态
		fmt.Printf("ld handler is processing req: %d\n", request.ID)
		// 如果出现异常，整个流程结束

	} else if ld.next != nil {
		ld.next.Handle(request)
	} else {
		return errors.New("该权重无节点可以处理")
	}
	return nil

}

func (ld *LDHandler) NextNode(node Handler) {
	ld.next = node
}

func TestResponsible(t *testing.T) {

	req1 := &Request{
		ID:     1,
		Weight: 4,
		MSG:    "leave 4 days",
	}
	req2 := &Request{
		ID:     2,
		Weight: 20,
		MSG:    "leave 20 days",
	}

	reqCh := make(chan *Request, 10)
	reqCh <- req1
	reqCh <- req2

	go func(reqCh chan *Request) {
		mt := new(MTHandler)
		ld := new(LDHandler)
		mt.NextNode(ld)
		for {
			select {
			case req := <-reqCh:
				mt.Handle(req)
			}
		}

	}(reqCh)

	time.Sleep(time.Second * 3)

}

func TestError(t *testing.T) {
	errors.New("error")
	_ = fmt.Errorf("this is an err code: %d, msg: %s", 404, "IER")
}
