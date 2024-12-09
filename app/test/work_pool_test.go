package test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

type Task func(ctx context.Context) error

type WorkPool struct {
	//如果这些变量只应该在WorkerPool结构体所属的包内部使用，就应该将它们定义为小写
	workerNum int
	taskQueue chan Task // 有缓冲的chan
	wg        sync.WaitGroup
}

func NewWorkPool(workerNum int) *WorkPool {
	return &WorkPool{
		workerNum: workerNum,
		taskQueue: make(chan Task, workerNum), // 个数如何确定比较好?
		// wg:        sync.WaitGroup{}, 不需要初始化，零值具备它的初始状态特征
	}
}

// exeTime 每个worker的执行耗时
func (wp *WorkPool) Start(exeTime time.Duration) error {
	// 开启workerNum个协程数 去监控队列
	for i := 0; i < wp.workerNum; i++ {
		wp.wg.Add(1)
		go func(workerId int) {
			defer wp.wg.Done()
			for task := range wp.taskQueue { // 阻塞，如果channel被关闭，就会退出
				ctx, _ := context.WithTimeout(context.Background(), exeTime)
				fmt.Printf("协程workerId=(%d)正在执行任务\n", workerId)
				err := task(ctx)
				if err != nil {
					fmt.Errorf("err : (%w)", err)
				}
				// 处理结果，如果处理可以提供管道对外暴露
			}
		}(i)
	}
	return nil
}

func (wp *WorkPool) AddTask(task Task) error {
	wp.taskQueue <- task // 阻塞
	return nil
}

func (wp *WorkPool) StopWorkerPool() error {
	close(wp.taskQueue)
	wp.wg.Wait()
	fmt.Println("工作池关闭成功")
	return nil
}

func TestWorkPool(t *testing.T) {

	tasks := []Task{
		func(ctx context.Context) error {
			select {
			case <-ctx.Done():
				fmt.Println("func 1 开始执行")
				fmt.Println("ctx 已经终止")
				return ctx.Err()
			default:
				time.Sleep(time.Millisecond * 5)
				for i := 0; i < 5; i++ {
					fmt.Println(i)
				}
				return nil
			}
		},
		func(ctx context.Context) error {
			select {
			case <-ctx.Done():
				fmt.Println("ctx 已经终止")
				return ctx.Err()
			default:
				fmt.Println("func 2 开始执行")
				time.Sleep(time.Millisecond * 5)
				for i := 5; i < 10; i++ {
					fmt.Println(i)
				}
				return nil
			}
		},
		func(ctx context.Context) error {
			select {
			case <-ctx.Done():
				fmt.Println("ctx 已经终止")
				return ctx.Err()
			default:
				fmt.Println("func 3 开始执行")
				time.Sleep(time.Millisecond * 5)
				for i := 10; i < 15; i++ {
					fmt.Println(i)
				}
				return nil
			}
		},
	}
	wp := NewWorkPool(2)
	wp.Start(time.Second * 30)

	for _, task := range tasks {
		wp.AddTask(task)
	}

	wp.StopWorkerPool()
}
