/*

 */
package main

import (
	"context"
	"fmt"
	"time"
)

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
}

func fatherGoroutine() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go getNum(ctx, 6)
	go getNum(ctx, 23)
	getNum(ctx, 2)
	time.Sleep(4 * time.Second)
	return
}

func main() {
	fatherGoroutine()

	time.Sleep(10 * time.Second)
}

/*
start: 2
start: 6
start: 23
currNum: 6
done: 23
done: 2
*/
