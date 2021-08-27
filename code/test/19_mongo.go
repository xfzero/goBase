package main

//使用原生 go get gopkg.in/mgo.v2

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

type Student3 struct {
	Name   string `bson: "name"`
	Age    int    `bson: "age"`
	Sid    string `bson: "sid"`
	Status int    `bson: "status"`
}

func main() {
	mongo, err := mgo.Dial("127.0.0.1") // 建立连接

	defer mongo.Close()

	if err != nil {
		fmt.Println("链接失败")
		return
	}

	client := mongo.DB("test").C("student") //选择数据库和集合

	//创建数据
	data := Student3{
		Name:   "jike",
		Age:    18,
		Sid:    "s1000233",
		Status: 1,
	}

	//插入数据
	cErr := client.Insert(&data)

	if cErr != nil {
		fmt.Println(cErr)
	}
	fmt.Println("success")
}
