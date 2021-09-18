/*
Mutex vs Channel:
通常，当goroutine需要相互通信时使用通道，当确保同一时间只有一个goroutine能访问代码的关键部分时使用互斥锁。
*/
package main

import (
	"fmt"
	"sync"
)

//使用WaitGroupi
func getNumber1() int {
	var i int

	//初始化一个waitGroup
	var wg sync.WaitGroup

	//Add(1) 通知程序有一个需要等待完成的任务
	wg.Add(1)

	go func() {
		i = 1
		//表示正在等待的程序已经执行完成
		wg.Done()
	}()

	//阻塞当前程序直到等待的程序都执行完成为止
	wg.Wait()

	return i
}

//使用通道阻塞
func getNumber2() int {
	var i int

	//创建一个通道，在等待的任务完成时会想通道发送一个空结构体
	done := make(chan struct{})

	go func() {
		i = 2
		//执行完成后向通道发送一个空结构体
		done <- struct{}{}
	}()

	//从程序接收值将会组赛程序，直到有值发送给done通道为止
	<-done
	return i
}

//使用Mutex
func getNumber3() {

}

func main() {
	num1 := getNumber1()
	fmt.Println(num1)

	num2 := getNumber2()
	fmt.Println(num2)
}
