package main

import "fmt"

//golang非面向对象语言，可以使用struct实现面向对象的特性

type Duck struct {
	Call string
	Color string
}

//实现了继承
type LittleDuck struct {
	Duck
}

func main() {
	ld := &LittleDuck{}
	ld.Color = "yel"

	fmt.Println(ld.Color)
}