package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"slices"
	"sync"
	"testing"
	"time"
	"trpc.group/trpc-go/trpc-go/log"
)

/**
你需要实现一个并发的任务调度器，能够处理多个并发任务并限制同时执行的任务数量。
每个任务是一个函数，任务执行可能需要一段时间。
调度器的目的是控制同时执行的任务数量（即并发度），以防止任务数过多导致系统资源过载。
*/

type TaskDispatcher struct {
	maxTask int
	c       chan func()
	wg      *sync.WaitGroup // 指针不进行初始化是nil
}

func NewTaskDispatcher(maxTaskNum int, chanCap int) *TaskDispatcher {
	return &TaskDispatcher{
		maxTask: maxTaskNum,
		c:       make(chan func(), chanCap),
		wg:      new(sync.WaitGroup),
	}
}

func (td *TaskDispatcher) AddNewTask(f func()) bool {
	select {
	case td.c <- f:
		td.wg.Add(1) // 这里加，消费完毕减少，这样是否合理
		return true
	case <-time.After(3 * time.Second):
		return false
	}
}

func (td *TaskDispatcher) RunDispatcher() {
	for i := 0; i < td.maxTask; i++ {
		go func(i int) {
			for f := range td.c {
				log.Info(fmt.Sprintf("当前第%d个协程正在执行任务", i))

				if f != nil {
					f()
				}

				time.Sleep(time.Second * time.Duration(rand.Intn(3)+1))
				fmt.Println("当前函数执行完毕")
				td.wg.Done()
			}
		}(i)
	}
}

func (td *TaskDispatcher) StopDispatcher() {

	defer close(td.c)
	td.wg.Wait()
}

func TestTaskDispatcher(t *testing.T) {
	td := NewTaskDispatcher(3, 10)

	td.RunDispatcher() //channel先有接收方，才向channel中发数据

	for i := 0; i < 20; i++ {
		td.AddNewTask(func() {
			fmt.Printf("这里是任务%v\n", i)
		})
	}

	td.StopDispatcher()
}

func TestAlgo3(t *testing.T) {
	test1 := []int{
		100, 4, 200, 1, 3, 2,
	}
	test2 := []int{0, 3, 7, 2, 5, 8, 4, 4, 4, 6, 0, 1}
	//res1 := continueNum(test1)
	//res2 := continueNum(test2)
	res1 := findContinueNum2(test1)
	res2 := findContinueNum2(test2)
	assert.Equal(t, 4, res1)
	assert.Equal(t, 9, res2)
}

func continueNum(nums []int) int {
	// 应该去重一下
	slices.Sort(nums)
	fmt.Println(nums)
	res := 0
	curRes := 1
	for i, val := range nums {

		if i+1 < len(nums) {
			if val+1 == nums[i+1] {
				curRes++
				res = max(curRes, res)
			} else {
				curRes = 1
			}
		}
	}
	return res
}

func findContinueNum2(nums []int) int {
	m := map[int]struct{}{}

	for _, val := range nums {
		m[val] = struct{}{}
	}

	res := 1
	for k, _ := range m {
		curRes := 1
		// 每个key验证
		if _, exist := m[k-1]; !exist {
			// 第一个数，统计最长
			num := k + 1
			for _, ok := m[num]; ok; {
				curRes++
				res = max(res, curRes)
				num++
				_, ok = m[num]
			}
		}
	}
	return res
}
