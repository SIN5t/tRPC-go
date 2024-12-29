package test

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

/*
*

Future模式是一种并发编程模式，它允许程序在等待某个操作完成的同时继续执行其他任务。
在Future模式中，一个操作（通常是耗时的）被提交执行，并立即返回一个future对象，这个对象代表了操作的结果。
其他部分的程序可以继续执行，而不必等待操作完成。
当需要操作的结果时，程序可以检查future对象，如果操作已完成，则直接获取结果；如果操作尚未完成，则可以阻塞等待结果。

 1. 将请求丢进一个结构体，开始异步地执行这耗时请求
 2. 主函数开始处理其他任务，不用等上一个执行完毕
 3. 主函数其他逻辑处理完毕之后，阻塞等待之前的处理结果
*/
type Call struct {
	Request  interface{}
	Response interface{}
	Done     chan *Call
}

// 给一个函数，丢进Call的成员，返回一个Call的实例
func foo(request interface{}, response interface{}, done chan *Call, runTask func(c *Call), wg *sync.WaitGroup) (*Call, error) {
	if done == nil {
		done = make(chan *Call, 10)
	} else if cap(done) == 0 { // 无缓冲的channel cap = 0
		return nil, errors.New("cap of done chan is 0")
	}

	call := new(Call)
	call.Request = request
	call.Response = response
	call.Done = done

	go func(c *Call) {
		runTask(c)
	}(call)

	return call, nil
}

func runTask(c *Call) {
	req := c.Request
	time.Sleep(time.Second * 2)
	resp := c.Response.(string) + "hello"
	c.Done <- &Call{
		Request:  req,
		Response: resp,
	}
	close(c.Done)
}

func TestSimpleFuture(t *testing.T) {
	var wg sync.WaitGroup
	call, err := foo("req", "resp", nil, runTask, &wg)
	if err != nil {
		fmt.Errorf("err future (%w)", err)
	}

	// 执行耗时操作
	time.Sleep(time.Millisecond * 10)
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Int31())
	}

	// 获取之前的执行结果
	resCall := <-call.Done
	fmt.Println(resCall.Response)

}
