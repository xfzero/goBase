package main

import "fmt"

func Defer() string{
	a := 2

	if a == 1 {
		fmt.Println("a:1")
		return "1"
	}
	defer fmt.Println("defer")

	if a == 2 {
		fmt.Println("a:2")
		return "2"
	}
	return "9"
}
/*
golang中defer,panic,recover是很常用的三个特性，三者一起使用可以充当其他语言中try…catch…的角色，
而defer本身又像其他语言的析构函数
*/

func main() {
	d := Defer()
	fmt.Println("d:",d)

	//a=1 a:1 -> d:1

	//a=2 a:2 -> defer -> d:2
}