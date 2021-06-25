package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

func httpRouter() {
	http.HandleFunc("/",RouterHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func RouterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	action := r.Form.Get("action")

	fmt.Println("request action:",action)

	data := HandleRequest()

	HttpOutput(w, 1, "success",data)
}

/*func HandleRequest() interface{} {
	arr := [...]int{1,2,4,5}
	return arr
}*/

func HandleRequest() [4]int {
	arr := [...]int{1,2,4,5}
	return arr
}

func HttpOutput(w http.ResponseWriter, ret, desc, data interface{}) {
	//fmt.Fprintf(w, "请求成功")

	var result = make(map[string]interface{})

	result["ret"] = ret
	result["desc"]= desc
	result["data"]= data

	show_result, _ := json.Marshal(result)

	w.Write(show_result) 
}

//启动服务：go run 08_http.go，请求：http://localhost:8000/?action=c
func main() {
	httpRouter()
}
