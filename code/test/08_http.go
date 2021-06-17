package main

import (
	"fmt"
	"log"
	"net/http"
)

func httpRouter() {
	http.HandleFunc("/",RouterHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func RouterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	action := r.Form.Get("action")

	fmt.Println(action)

	fmt.Fprintf(w, "请求成功")
}

//启动服务：go run 08_http.go，请求：http://localhost:8000/?action=c
func main() {
	httpRouter()
}

