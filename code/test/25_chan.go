package main

import (
	"fmt"
	"time"
)

type GrpcTask struct {
	MsgChan chan int
}

func main() {
	m := &GrpcTask{
		MsgChan: make(chan int, 2),
	}

	go func() {
		for msg := range m.MsgChan {
			fmt.Println("request:", msg)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			m.MsgChan <- i
			fmt.Println("msg:", i)
		}
	}()

	time.Sleep(15 * time.Second)
}
