package test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
多线程打印1-20并要求顺序
*/

func Print(ctx context.Context, ch chan int, id int) error {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctx 超时")
			return ctx.Err()

		default:
			num := <-ch
			fmt.Printf("worker: %d, print num: %d\n", id, num)
			if num == 20 {
				close(ch)
				return nil
			}
			ch <- num + 1
		}
	}

}

func TestPrintNum(t *testing.T) {
	ch := make(chan int)

	ctx, cancelFun := context.WithTimeout(context.Background(), time.Second*3)
	defer func() {
		cancelFun()
		close(ch)
	}()
	for i := 0; i < 4; i++ {
		go Print(ctx, ch, i)
	}
	ch <- 1
	select {
	case <-ctx.Done():
		return
	case <-time.After(time.Second * 5):
		return
	}

}

// 分别打印奇数偶数

func TestAnotherSolution(t *testing.T) {
	//仅仅把无缓冲的channel当成一把锁使用

	wg := sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan int)

	go func() {
		defer wg.Done()
		// 如果channel中没有数据，<-ch阻塞，直到channel中有一个数据时，取出channel的数据并判断i的值是否是偶数，如果是偶数，则打印
		for i := 1; i < 10; i += 1 {
			<-ch
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
	}()
	// 第一次循环时，channel为空，可以向channel写入数据，此时i=1，所以，可以打印出1
	// 第二次循环时，channel不为空，所以 ch <- 0语句阻塞，直到channel内没有数据
	go func() {
		defer wg.Done()
		for i := 1; i < 10; i += 1 {
			ch <- 0
			if i%2 == 1 {
				fmt.Println(i)
			}
		}
	}()

	wg.Wait()

}
