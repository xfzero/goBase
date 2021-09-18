package main

import (
	"fmt"
)

type People interface {
	ReturnName() string
	ReturnAge() uint32
}

type Student struct {
	Name string
	Age  uint32
}

// 定义结构体的一个方法。
// 突然发现这个方法同接口People的所有方法(就一个)，此时可直接认为结构体Student实现了接口People
func (this Student) ReturnName() string {
	return this.Name
}

func (this Student) ReturnAge() uint32 {
	return this.Age
}

func (this *Student) getId() {
	fmt.Println(this.Id)
}

func main() {
	cbs := Student{Name: "tom", Age: 18}

	var p People

	// 因为Students实现了接口所以直接赋值没问题
	// 如果没实现会报错：cannot use cbs (type Student) as type People in assignment:Student does not implement People (missing ReturnName method)
	p = cbs

	name := p.ReturnName()

	fmt.Println(name)
}
