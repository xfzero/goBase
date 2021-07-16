package main

import (
	"fmt"
	"net/http"
	"time"
)

func fetch(ch chan<- string,i int) {
	url := "https://www.mlytics.com/"

	res := "get url end" + fmt.Sprintf("%d", i)

	for i := 0; i < 20; i++ {
		_, err := http.Get(url)
		if err != nil {
	        
		}
	}
	
	//fmt.Println("fetch end")
	ch <- res
}

func main() {
	fmt.Println("main start:",time.Now())
	ch := make(chan string)
	for i := 0; i < 5; i++ {
		go fetch(ch,i)
	}

	fmt.Println("all url start:",time.Now())

	for i := 0; i < 5; i++ {
		// 注释后，程序不等待,5改为6 会一直等待，改为4，等待任意4个goroutine执行结束
	    fmt.Println(<-ch)
	}

	fmt.Println("all url end:",time.Now())
}