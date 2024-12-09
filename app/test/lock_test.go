package test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
*
对变量执行2000次+1操作
5个协程并发执行,且支持通过 context 取消操作
*/
type num struct {
	num int
	sync.Mutex
	sync.WaitGroup
}

func AddWithLock(ctx context.Context, n *num) {
	for i := 0; i < 2000; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("ctx time out")
			n.Done()
			return
		default:
			n.Lock()
			n.num++
			n.Unlock()
		}
	}
	n.Done()
}
func TestLock(t *testing.T) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	n := &num{
		num:       0,
		Mutex:     sync.Mutex{},
		WaitGroup: sync.WaitGroup{},
	}
	n.WaitGroup.Add(5)

	defer cancelFunc()
	for i := 0; i < 5; i++ {
		go AddWithLock(ctx, n)
	}
	n.Wait()

	fmt.Println(n.num)
}
