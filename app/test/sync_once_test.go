package test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

type Num struct {
	n int
}

type singletonInstance struct {
	n int
}

var singleInstance *singletonInstance

func TestSyncOnce(t *testing.T) {

	num := &Num{n: 10}

	var add = func() {
		num.n++
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			return
		default:
			once := sync.Once{}
			for i := 0; i < 10; i++ {
				once.Do(add)
			}
		}
	}(ctx)
	time.Sleep(time.Second)
	fmt.Println(num)

	select {
	case <-ctx.Done():
		fmt.Println("time out")
	}

}

func Singleton() *singletonInstance {
	if singleInstance != nil {
		return singleInstance
	}

	mutex := &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	return &singletonInstance{n: 10}
}

func TestSingleton(t *testing.T) {
	singleton := Singleton()
	fmt.Println(singleton)

}
