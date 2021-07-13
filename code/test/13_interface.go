package main

import "fmt"

//一个类型实现多个接口

type People interface {
	ReturnName() string
}

type Role interface {
	ReturnRole() string
}

type Student struct {
	Name string
}

func(s Student) ReturnName() string {
	return s.Name
}

func(s Student) ReturnRole() string {
	return "学生"
}

func main() {
	cbs := Student{Name: "tony"}

	var p People
	var r Role

	p = cbs //由于Student实现了People所有方法，接口实现成功，可直接赋值
	r = cbs //由于Student实现了Role所有方法，接口实现成功，可直接赋值

	name := p.ReturnName()
	fmt.Println(name)

	role := r.ReturnRole()
	fmt.Println(role)
}