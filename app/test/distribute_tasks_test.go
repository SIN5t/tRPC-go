package test

import (
	"fmt"
	"sync"
	"testing"
)

/**
计算100个自然数的和。将计算任务拆分为多个task，每个task启动一个goroutine进行处理。代码如下：
*/

type AddTask struct {
	Begin int
	End   int
	Res   chan int
}

func (addTask *AddTask) Do() {
	res := 0
	for i := addTask.Begin; i < addTask.End; i++ {
		res += i
	}
	addTask.Res <- res
}

func TestDistribute(t *testing.T) {
	numOfTask := 5
	resChan := make(chan int, numOfTask)
	addTaskChan := make(chan *AddTask, numOfTask)

	go initTask(addTaskChan, numOfTask, 11, resChan)
	go DistributeTasks(addTaskChan, resChan)

	res := ProcessRes(resChan)
	fmt.Println(res)

}

func ProcessRes(resChan <-chan int) (res int) {
	for r := range resChan {
		res += r
	}
	return res
}

func initTask(tasksChan chan *AddTask, numOfTask int, num int, resChan chan int) {
	n := num / numOfTask
	m := num % numOfTask

	defer close(tasksChan)
	for i := 0; i < n; i++ {
		b := numOfTask * i
		e := numOfTask * (i + 1)
		task := &AddTask{
			Begin: b,
			End:   e,
			Res:   resChan,
		}
		tasksChan <- task
	}
	if m != 0 {
		tasksChan <- &AddTask{
			Begin: n * numOfTask,
			End:   num + 1,
			Res:   resChan,
		}
	}
}

func DistributeTasks(taskChan <-chan *AddTask, resChan chan int) {
	var wg sync.WaitGroup
	for task := range taskChan {
		wg.Add(1)
		go func(t *AddTask) {
			defer wg.Done()
			t.Do()
		}(task) // 注意要当作参数传入，而不是直接在 开启的协程 内部调用task，
	}
	wg.Wait()
	close(resChan)
}
