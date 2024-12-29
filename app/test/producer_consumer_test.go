package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type ConsumeTask struct {
	id       int
	execFunc func()
}

type Producer struct {
	id int
	ch chan *ConsumeTask
}

type Consumer struct {
	id int
	ch chan *ConsumeTask
}

func TestProducerConsumer(t *testing.T) {

	taskCh := make(chan *ConsumeTask, 100)
	defer close(taskCh)

	producers := make([]*Producer, 0, 10)
	consumers := make([]*Consumer, 0, 5)
	// 创建十个消费者
	for i := 0; i < 10; i++ {
		consumers = append(consumers, NewConsumer(i, taskCh))
	}
	// 创建十个生产者
	for i := 0; i < 10; i++ {
		producers = append(producers, NewProducer(i, taskCh))
	}

	tasks := make(chan *ConsumeTask, 100)
	for i := 0; i < 100; i++ {
		t := &ConsumeTask{
			id: i,
			execFunc: func() {
				dura := time.Duration(rand.Intn(3)+1) * time.Second
				time.Sleep(dura)
				fmt.Println("已经睡眠:", dura.String(), ", 任务", i, "执行中...")
			},
		}
		tasks <- t
	}
	defer close(tasks)

	for _, producer := range producers {
		// 闭包问题
		//go func() {
		//	for t := range tasks {
		//		producer.Produce(t)
		//	}
		//}()

		go func(p *Producer) {
			for t := range tasks {
				p.Produce(t)
			}
		}(producer)
	}

	for _, consumer := range consumers {
		//// 闭包问题
		//go func() {
		//	consumer.Consume()
		//}()

		go func(c *Consumer) {
			c.Consume()
		}(consumer)
	}

	time.Sleep(time.Second * 10)

}

func NewProducer(id int, ch chan *ConsumeTask) *Producer {
	return &Producer{
		id: id,
		ch: ch,
	}
}

func NewConsumer(id int, ch chan *ConsumeTask) *Consumer {
	return &Consumer{
		id: id,
		ch: ch,
	}
}

func (p *Producer) Produce(task *ConsumeTask) {
	p.ch <- task
}

func (c *Consumer) Consume() {
	for task := range c.ch {
		fmt.Printf("消费者%d,正在消费\n", c.id)
		task.execFunc()
	}
}
