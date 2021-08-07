package main

import "fmt"

func test1() {
	x := 1
	p := &x 		//p是整形指针，指向x
	fmt.Println(*p) //1，p指向的变量写成*p 表达式*p获取变量的值
	*p = 2 			//等于x=2
	fmt.Println(x) 	//2
}

func test2() {
	var name *string
	dog := "tony"
	name = &dog

	fmt.Println(dog) //tony
	fmt.Println(&dog) //0xc0000821e0
	fmt.Println(name) //0xc0000821e0
	fmt.Println(*name) //tony

	fmt.Printf("%s\n",name)
}

func test3() {//两个指针当切仅当指向同一个变量或者都是nil的情况才相等
	var x, y int
	fmt.Println(&x == &x, &x == &y, &x == nil)// true false false

	var z *string
	fmt.Println(z)
}

func f() *int {
	v := 1
	return &v
}

func main() {
	// test1()

	// test2()

	// test3()

	var p = f()
	fmt.Println(p, *p)
	fmt.Println(f(), f())

}