package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCancelContext(t *testing.T) {

	ctx, cancelFunc := context.WithCancel(context.Background())

	for i := 0; i < 10; i++ {
		go worker(ctx, i)
	}
	select {
	case <-ctx.Done():
		fmt.Println("main ctx canceled")
	case <-time.After(1 * time.Second):
		cancelFunc() // 主函数开始取消

	}
}

func worker(ctx context.Context, id int) {

	fmt.Printf("worker %v is working\n", id)

	select {
	case <-ctx.Done():
		fmt.Println("ctx is done")
	case <-time.After(1 * time.Second):
		fmt.Printf("worker %v finishde the work \n", id)
	}

}
