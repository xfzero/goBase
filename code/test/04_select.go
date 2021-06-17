package main

import (
	"fmt"
	"time"
)

//chs := make(chan int,1)//自动推导类型不能在方法体外使用？
var chs chan int = make(chan int,1)

func write() {
	time.Sleep(3*time.Second)
	chs <-88
}

func read() {
	select {
	case ch1 :=<-chs:
		fmt.Println(ch1)
		return
	// case <-time.After(time.Second):
	case <-time.After(4*time.Second):
		fmt.Println("read time out")
		return
	}
}

func main() {
	go write()
	read()
}