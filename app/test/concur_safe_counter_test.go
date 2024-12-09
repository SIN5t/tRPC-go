package test

import (
	"fmt"
	"sync"
	"testing"
)

/**
实现一个并发安全的计数器
要求：
	使用 Go 的并发机制（如 goroutines 和 channels）实现一个计数器。
	计数器可以在多个 goroutine 中并发访问。
	实现一个 Counter 结构体，并提供如下方法：
	Increment(): 增加计数器的值。
	Get() int: 返回计数器当前的值。
	确保对计数器的访问是并发安全的，避免数据竞争。
*/

type Counter struct {
	num int
	mu  sync.RWMutex
}

func (c *Counter) Increment() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.num++
	return nil
}

func (c *Counter) Get() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.num
}

func TestCounter(t *testing.T) {
	wg := sync.WaitGroup{}
	counter := &Counter{
		num: 0,
	}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(c *Counter) {
			c.Increment()
			fmt.Println(c.Get())
			wg.Done()
		}(counter)
	}
	wg.Wait()

}
