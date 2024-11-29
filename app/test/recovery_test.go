package test

import (
	"fmt"
	"testing"
)

func recoveryDeferPanicOrder1() {

	fmt.Println("1. 程序开始")
	defer fmt.Println("2. defer fmt print")

	// 模拟致命错误
	panic("fatal err, panic !")
	defer fmt.Println("panic后这行代码执行不到")
}

func recoveryDeferPanicOrder2() {

	fmt.Println("1. 程序开始")
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("inner函数中的defer通过recover捕获到异常，异常信息: %v\n", err)
		}

		fmt.Println("2. defer fmt print")
	}()

	// 模拟致命错误
	panic("fatal err, panic !")
	defer fmt.Println("after panic line")
}

func function1() {
	recoveryDeferPanicOrder2()
	fmt.Println("上面发生异常，被捕捉后，这里继续执行")
}

func TestRecovery(t *testing.T) {
	//recoveryDeferPanicOrder2()
	//recoveryDeferPanicOrder1()
	function1()

	fmt.Println("hello")

}
