package test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/**
有三个需要并发执行的函数，三个函数会返回同一类型的值，
且三个函数执行时间在2ms到10ms之间不等，而主程序要求在5ms内返回结果，
若5ms内没有执行完毕，则强制返回结果，这个时候某个函数可能还没有返回因此没有值。

func searchText(query string) ResultText
func searchImage(query string) ResultText
func searchVideo(query string) ResultText
*/

type ResultText string
type ResultImage string
type ResultVideo string
type AllRes struct {
	ResultText
	ResultImage
	ResultVideo
}

func TestConcur(t *testing.T) {
	ctx, cancelFun := context.WithTimeout(context.Background(), time.Millisecond*5)
	defer cancelFun()
	textCh := make(chan ResultText)
	imageCh := make(chan ResultImage)
	videoCh := make(chan ResultVideo)
	go searchText(ctx, "text", textCh)
	go searchImage(ctx, "image", imageCh)
	go searchVideo(ctx, "video", videoCh)
	res := AllRes{}

	for {
		select {
		case textRes := <-textCh:
			res.ResultText = textRes
		case imageRes := <-imageCh:
			res.ResultImage = imageRes
		case videoRes := <-videoCh:
			res.ResultVideo = videoRes

		case <-ctx.Done():
			fmt.Println(res)
			return // 实际业务中，应该在这里返回数据
		}
	}

}

func searchText(ctx context.Context, query string, ch chan ResultText) ResultText {

	res := ResultText("")
	select {
	case <-ctx.Done():
		fmt.Errorf("超时")
		ch <- ""
		return res
	default:
		// 耗时操作,业务逻辑，实际业务逻辑应该封装为一个函数，不写在这个部分，这个部分主要做超时控制
		duration := time.Duration(rand.Intn(9) + 2)
		time.Sleep(time.Millisecond * duration)
		res = ResultText(query + duration.String())
		ch <- res
	}

	return res
}

func searchImage(ctx context.Context, query string, ch chan ResultImage) ResultImage {

	res := ResultImage("")
	select {
	case <-ctx.Done():
		fmt.Errorf("超时")
		ch <- ""
		return res
	default:
		// 耗时操作,业务逻辑，实际业务逻辑应该封装为一个函数，不写在这个部分，这个部分主要做超时控制
		duration := time.Duration(rand.Intn(9) + 2)
		time.Sleep(time.Millisecond * duration)
		res = ResultImage(query + duration.String())
		ch <- res
	}

	return res
}
func searchVideo(ctx context.Context, query string, ch chan ResultVideo) ResultVideo {

	res := ResultVideo("")
	select {
	case <-ctx.Done():
		fmt.Errorf("超时")
		ch <- ""
		return res
	default:
		// 耗时操作,业务逻辑，实际业务逻辑应该封装为一个函数，不写在这个部分，这个部分主要做超时控制
		duration := time.Duration(rand.Intn(9) + 2)
		time.Sleep(time.Millisecond * duration)
		res = ResultVideo(query + duration.String())
		ch <- res
	}
	return res
}
