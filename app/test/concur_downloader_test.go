package test

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

/*
*
你需要编写一个模拟并发下载文件的程序。假设有多个文件需要下载，下载任务是独立的，但是每次下载一个文件时，
它需要消耗一定的时间。为了加快下载速度，你应该使用 Goroutines 来实现并发下载。
*/
type ConcurDownloader struct {
	DownLoadUrl  string
	DownLoadDest string
}

func (c ConcurDownloader) NewDownloader(fromUrl string, dest string) *ConcurDownloader {
	return &ConcurDownloader{
		DownLoadUrl:  fromUrl,
		DownLoadDest: dest,
	}
}

func concurDownLodeFile(ctx context.Context, c []*ConcurDownloader, ch chan<- string) {
	defer close(ch)
	wg := new(sync.WaitGroup)
	for i, task := range c {

		wg.Add(1)
		go func(c *ConcurDownloader, ctx context.Context, i int) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				fmt.Println("ctx 超时，子协程：", i, "主动结束")
				return
			default:
				fmt.Println(c.DownLoadUrl, "：-》已下载")                      // 模拟下载
				randomTime := time.Duration(rand.Intn(4)+2) * time.Second // 1-2s的随机时间
				time.Sleep(randomTime)                                    // 模拟存储
				//fmt.Println(c.DownLoadDest, "目的地址已经已获取")
				ch <- strconv.Itoa(i+1) + "协程任务完成" // 当前任务完成

			}
		}(task, ctx, i)
	}
	wg.Wait()

}

func TestConcurDownload(t *testing.T) {

	concurDownloaders := []*ConcurDownloader{
		&ConcurDownloader{
			DownLoadUrl:  "源url1",
			DownLoadDest: "下载地址1",
		},
		&ConcurDownloader{
			DownLoadUrl:  "源url2",
			DownLoadDest: "下载地址2",
		},
		&ConcurDownloader{
			DownLoadUrl:  "源url3",
			DownLoadDest: "下载地址3",
		},
	}
	ctx, cancelFun := context.WithTimeout(context.Background(), time.Second*3)
	defer cancelFun()
	ch := make(chan string, len(concurDownloaders))
	concurDownLodeFile(ctx, concurDownloaders, ch)

Label:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctx 超时")
			return
		//case info := <-ch: // 注意，如果channel关闭了还读，那就会一直读出空的值，for循环一直打印出空值！
		//
		//	fmt.Println(info)
		//}
		case info, ok := <-ch: // 或者用range写法
			if !ok {
				fmt.Println("所有任务已完成")
				break Label
			}
			fmt.Println(info)
		}
	}
	fmt.Println("主函数即将执行完毕")
}
