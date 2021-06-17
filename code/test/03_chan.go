package main

import "fmt"

var message2 chan string

func main() {
	message := make(chan string)
	go func() {
		message <- "ping"
	}()
	msg := <-message
	fmt.Println(msg)

	message2 = make(chan string)
	go ping()
	msg2 := <-message2
	fmt.Println(msg2)
}

func ping(){
	message2 <- "ping2"
}