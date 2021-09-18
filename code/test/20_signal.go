package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//创建一个os.signal channel
	sigs := make(chan os.Signal, 1)
	//创建一个bool channel
	done := make(chan bool, 1)

	//注册要接收的信号，syscall.SIGINT:接收ctrl+c,syscall.SIGTERM:程序退出
	//信号没有信号参数表示接收所有的信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	//此goroutine为执行组赛接受信号，一旦有了它，他就会打印出来
	//然后通知程序可以完成
	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	//不容许继续往sigs中存入内容,仅a输出awaiting signal。不会输出signal.Stop(sigs)之后的内容。
	//signal.Stop(sigs)

	//程序将在此处等待，直到它预期信号（如Goroutine所示）
	//在”done“上发送一个值,然后退出
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}

/*
执行 go run 20_signal.go 再敲入 ctrl+c，程序会输出

awaiting signal

interrupt
exiting

如果用kill pid 结束程序，则输出

awaiting signal
Terminated
*/
