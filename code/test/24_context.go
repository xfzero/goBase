/*
1. Context作用：
通过Context可以区分不同的goroutine请求，因为在Golang Severs中，每个请求都是在单个goroutine中完成的。

相互调用的goroutine之间通过传递context变量保持关联，这样在不用暴露各goroutine内部实现细节的前提下，
有效地控制各goroutine的运行。

通过传递Context就可以追踪goroutine调用树，并在这些调用树之间传递通知和元数据。

2. context常用的使用：
2.1) web编程中，一个请求对应多个goroutine之间的数据交互
2.2) 超时控制
2.3) 上下文控制

3. Context结构体：
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
3.1 Done:
Done方法在Context被取消或超时时返回一个close的channel,close的channel可以作为广播通知，告诉给context相关的函数要停止当前工作然后返回。

当一个父operation启动一个goroutine用于子operation，这些子operation不能够取消父operation。下面描述的WithCancel函数提供一种方式可以取消新创建的Context.

Context可以安全的被多个goroutine使用。开发者可以把一个Context传递给任意多个goroutine然后cancel这个context的时候就能够通知到所有的goroutine。

Done方法返回一个只读的chan，类型为struct{}，我们在goroutine中，如果该方法返回的chan可以读取，
则意味着parent context已经发起了取消请求，我们通过Done方法收到这个信号后，就应该做清理操作，
然后退出goroutine，释放资源。


3.2 Err:
Err方法返回context为什么被取消。

3.3 Deadline:
Deadline返回context何时会超时。即获取设置的截止时间的意思。
第一个返回式是截止时间，到了这个时间点，Context会自动发起取消请求；第二个返回值ok==false时表示没有设置截止时间，如果需要取消的话，需要调用取消函数进行取消。

3.4 Value:
value方法获取该Context上绑定的值，是一个键值对，所以要通过一个Key才可以获取对应的值，这个值一般是线程安全的。

4. 继承的Context:
4.1：BackGround（顶层Context：Background）：
BackGound是所有Context的root，不能够被cancel。

该Context通常由接收request的第一个goroutine创建，它不能被取消、没有值、也没有过期时间，常作为处理request的顶层context存在。

4.2：下层Context：WithCancel/WithDeadline/WithTimeout
创建了根节点之后，接下来就是创建子孙节点。为了可以很好的控制子孙节点，Context包提供的创建方法均是带有第二返回值（CancelFunc类型）。
它相当于一个Hook，在子goroutine执行过程中，可以通过触发Hook来达到控制子goroutine的目的（通常是取消，即让其停下来）。
再配合Context提供的Done方法，子goroutine可以检查自身是否被父级节点Cancel。
select {
    case <-ctx.Done():
        // do some clean…
}

my注：取消操作时，父goroutine调用cancel,子goroutine通过ctx.Done()主动取消

// 带cancel返回值的Context，一旦cancel被调用，即取消该创建的context
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

// 带有效期cancel返回值的Context，即必须到达指定时间点调用的cancel方法才会被执行
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)

// 带超时时间cancel返回值的Context，类似Deadline，前者是时间点，后者为时间间隔
// 相当于WithDeadline(parent, time.Now().Add(timeout)).
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

WithDeadline和WithTimeout的不同：
WithDeadline第二个参数是时间点，WithTimeout是时间间隔

5. context的优缺点：
5.1 缺点：
不知道子goroutine执行进度的情况下直接取消，控制不够精确。
每一个相关函数都必须增加一个context.Context类型的参数，且作为第一个参数，这对无关代码完全是侵入式的。

5.2 优点：
当衍生的子孙goroutine比较多时，取消比较简单。

6. Context接口：
Context接口并不需要我们实现，Go内置已经帮我们实现了2个，我们代码中最开始都是以这两个内置的作为最顶层的partent context，衍生出更多的子Context。

6.1 Background：
主要用于main函数、初始化以及测试代码中，作为Context这个树结构的最顶层的Context，也就是根Context。

6.2 TODO：
它目前还不知道具体的使用场景，如果我们不知道该使用什么Context的时候，可以使用这个。
这个就是创建一个占位用的context，可能在写程序的过程中还不能确定后期这个context的作用，所以暂时用这个占位，而不是使用nil

他们两个本质上都是emptyCtx结构体类型，是一个不可取消，没有设置截止时间，没有携带任何值的Context。
*/

package main

import (
	"context"
	"fmt"
	"time"
)

var nums []uint32 = []uint32{3, 6, 1, 100, 50, 10, 35, 23, 2, 846, 345, 68, 30, 458}

func getNum1(ctx context.Context, num uint32) {
	fmt.Println("start:", num)
	var currNum uint32

	for _, v := range nums {
		//每一秒判断一次ctx是否被取消
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

func process(ctx context.Context) {
	session, ok := ctx.Value("session").(int)
	if !ok {
		fmt.Println("something wrong")
		return
	}

	if session != 1 {
		fmt.Println("session 位通过")
		return
	}

	traceId := ctx.Value("trace_id").(string)
	fmt.Println("traceId:", traceId, "-session:", session)
}

func getNum2(ctx context.Context, num uint32) {
	fmt.Println("start:", num)
	var currNum uint32

	if deadline, ok := ctx.Deadline(); ok { //设置了deadl
		fmt.Println("deadline set:", deadline)
		if time.Now().After(deadline) {
			fmt.Println(ctx.Err().Error())
			return
		}
	}

	for _, v := range nums {
		//每一秒判断一次ctx是否被取消
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("done:", num)
			fmt.Println("ctxerr:", ctx.Err())
			//return ctx.Err()
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

//测试WithCancel
func testWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go getNum1(ctx, 6)
	go getNum1(ctx, 2)
	go getNum1(ctx, 1)

	//3秒后执行cancel()，子goroutine通过ctx.Done()收到消息取消当前goroutine
	time.Sleep(3 * time.Second)
}

//测试WithDeadline
func testWithDeadline() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))

	go getNum1(ctx, 6)
	go getNum1(ctx, 2)
	go getNum1(ctx, 1)

	//5秒后执行cancel有效
	time.Sleep(5 * time.Second)
	fmt.Println("5秒后执行cancel")
	cancel()
}

//测试WithTimeout
func testWithTimeout() {
	//WithTimeout 等价于 WithDeadline(parent, time.Now().Add(timeout)).
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go getNum1(ctx, 6)
	go getNum1(ctx, 2)
	go getNum1(ctx, 1)

	//3秒后ctx超时会主动触发cancel，子goroutine通过ctx.Done()收到超时的消息取消当前goroutine，当然可以在超时之前触发cancel
	time.Sleep(6 * time.Second)
}

//测试WithValue
func testWithValue() {
	ctx := context.WithValue(context.Background(), "trace_id", "100001")

	ctx = context.WithValue(ctx, "session", 1)

	process(ctx)
}

//测试Err和Deadline
func testContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go getNum2(ctx, 100)
	go getNum2(ctx, 35)

	time.Sleep(6 * time.Second)
}

func main() {
	//go testWithCancel()

	//go testWithDeadline()

	//go testWithTimeout()

	//go testWithValue()

	//以下测试done、err、Deadline、value
	//done和value在之前已经使用过，这里主要测试另外两个

	go testContext()

	time.Sleep(10 * time.Second)
}
