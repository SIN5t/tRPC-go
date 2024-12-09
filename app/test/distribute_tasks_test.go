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
	wg := sync.WaitGroup{}
	numOfTask := 8
	resChan := make(chan int, numOfTask)
	addTaskChan := initTask(numOfTask, 100, resChan)
	DistributeTasks(addTaskChan, numOfTask, wg)
	res := 0
	wg.Wait()
	close(resChan)
	for resCh := range resChan {
		res += resCh
	}
	fmt.Println(res)

}

func initTask(numOfTask int, num int, resChan chan int) chan *AddTask {
	n := num / numOfTask
	m := num % numOfTask

	tasksChan := make(chan *AddTask, numOfTask)
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
			End:   num,
			Res:   resChan,
		}
	}
	return tasksChan
}

func DistributeTasks(taskChan chan *AddTask, numOfTasks int, wg sync.WaitGroup) {

	for task := range taskChan {
		wg.Add(1)
		go func() {
			defer wg.Done()
			task.Do()
		}()
	}

}
