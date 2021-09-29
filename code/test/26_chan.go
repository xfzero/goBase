package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	gGrpcTask *GrpcTaskM
	ChanSize  = 2
)

func init() {
	fmt.Println("init")
}

func httpRouter() {
	http.HandleFunc("/", RouterHandler)
}

func RouterHandler(w http.ResponseWriter, r *http.Request) {
	//模拟收到请求转发
	msg := &GrpcRequest{
		SessionId: 100000,
		Cmd:       "activity",
		ServerId:  100001,
		Auth:      "params",
	}
	gGrpcTask.SendMsg(msg)

	r.ParseForm()

	action := r.Form.Get("action")

	fmt.Println("request action:", action)

	data := HandleRequest()

	HttpOutput(w, 1, "success", data)
}

func HttpOutput(w http.ResponseWriter, ret, desc, data interface{}) {
	//fmt.Fprintf(w, "请求成功")

	var result = make(map[string]interface{})

	result["ret"] = ret
	result["desc"] = desc
	result["data"] = data

	show_result, _ := json.Marshal(result)

	w.Write(show_result)
}

func HandleRequest() [4]int {
	arr := [...]int{1, 2, 4, 5}
	return arr
}

type GrpcRequest struct {
	SessionId uint64
	Cmd       string
	ServerId  uint32
	Auth      string
}

type GrpcTaskM struct {
	MsgChan []chan *GrpcRequest
}

type GrpcTask struct {
	chanIndex int
	msgChan   chan *GrpcRequest
}

func NewGrpcTaskM() *GrpcTaskM {
	m := &GrpcTaskM{
		MsgChan: make([]chan *GrpcRequest, ChanSize),
	}

	for i := 0; i < ChanSize; i++ {
		m.MsgChan[i] = make(chan *GrpcRequest, 1024)
		m.CreateTask(i)
	}

	return m
}

func (this *GrpcTaskM) CreateTask(chanIndex int) {
	ret := NewGrpcTask(chanIndex, this.MsgChan[chanIndex])
	go ret.Process()
}

func NewGrpcTask(index int, ch chan *GrpcRequest) *GrpcTask {
	ret := &GrpcTask{
		chanIndex: index,
		msgChan:   ch,
	}

	return ret
}

func (this *GrpcTask) Process() {
	fmt.Println("process start")
	for msg := range this.msgChan {
		fmt.Println("sendGrpcRequest:", msg)
	}
	fmt.Println("process end")
}

func (this *GrpcTaskM) SendMsg(msg *GrpcRequest) {
	hashId := 1
	this.MsgChan[hashId] <- msg //写入消息后，Process处理
}

func grpcTask() {
	gGrpcTask = NewGrpcTaskM()
}

func main() {
	go func() {
		httpRouter()
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
	}()

	grpcTask()

	time.Sleep(20 * time.Second) //实际应用中可以用signal代替
}
