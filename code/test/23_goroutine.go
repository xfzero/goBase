/*
假设有一个超长的切片，切片的元素类型为int，切片中的元素为乱序排列。
限时5秒，使用多个goroutine查找切片中是否存在给定值，在找到目标值或者超时后立刻结束所有goroutine的执行。

比如切片为：[23, 32, 78, 43, 76, 65, 345, 762, …… 915, 86]，查找的目标值为345，
如果切片中存在目标值程序输出:"Found it!"并且立即取消仍在执行查找任务的goroutine。
如果在超时时间未找到目标值程序输出:"Timeout! Not Found"，同时立即取消仍在执行查找任务的goroutine。
*/
package main

import (
	"context"
	"fmt"
	"time"
)

var done chan bool

var nums []uint32 = []uint32{3, 6, 1, 100, 50, 10, 35, 23, 2, 846, 345, 68, 30, 458}

func getNum(ctx context.Context, num uint32) {
	fmt.Println("start:", num)
	var currNum uint32

	for _, v := range nums {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("done:", num)
			return
		default:
		}

		if v == num {
			currNum = num
			break
		}
	}

	fmt.Println("currNum:", currNum)
	done <- true //找打值后向done写入消息
}

func fatherGoroutine() {
	done = make(chan bool, 1)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go getNum(ctx, 6)
	go getNum(ctx, 10)
	go getNum(ctx, 2)

	//先找到6后done会接收到消息->执行cancel()取消其他的两个goroutine
	<-done
	return
}

func main() {
	fatherGoroutine()

	time.Sleep(30 * time.Second)
}

/*
start: 6
start: 2
start: 10
currNum: 6
done: 10
done: 2

当任意一个getNum找到值后，会向done通道写入消息
fatherGoroutine从done读取到消息后认为至少有一个goroutine执行成功->执行cancel方法取消所有的goroutine
如果所有goroutine在10秒钟类没有执行成功，则超时结束
*/
