package main

import "fmt"

type People interface {
	ReturnName() string
}

type Student struct {
	Name string
}

type Teacher struct {
	Name string
}

func (s Student) ReturnName() string {
	return s.Name
}

func (t *Teacher) ReturnName() string {
	return t.Name
}

func main() {
	cbs := Student{Name: "tony"}
	sss := Teacher{Name: "tom"}

	//值类型
	var a People
	a = cbs
	name := a.ReturnName()
	fmt.Println(name)

	// 指针类型
	a = &sss //由于是指针类型，所以赋值时要加上&
	name = a.ReturnName()
	fmt.Println(name)
	//不能使用 a=sss 因为是Teacher的指针实现了ReturnName方法，Teacher本身没实现
}