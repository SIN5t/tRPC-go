package test

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

type mutexCh struct { //应该大写，给外部所有包用
	ch chan int // 私有变量，否则外部操作
}

func NewMutexCh() *mutexCh {
	return &mutexCh{ch: make(chan int, 1)}
}

func (mc *mutexCh) TryLock() {
	mc.ch <- 1
}

func (mc *mutexCh) TryUnlock() {
	select {
	case <-mc.ch:
	default:
		panic("unlock an unlocked lock")
	}
}

func TestChLock(t *testing.T) {
	mc := NewMutexCh()
	m := map[int]string{
		0: "hello",
		1: "world",
	}
	for i := 0; i < 10; i++ {
		go func(mc *mutexCh, m map[int]string, num int) {
			mc.TryLock() //阻塞获取锁
			m[0] = m[0] + strconv.Itoa(i)
			mc.TryUnlock()

		}(mc, m, i)
	}

	select {
	default:
		<-time.After(time.Second)
		fmt.Println(m)
	}

}
