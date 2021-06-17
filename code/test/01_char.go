package main

import "fmt"

func main() {
	c := 'a'
	fmt.Println(c) //97 字符类型是特殊的整形
	fmt.Printf("%c\n",c) //a

	s := fmt.Sprintf("%c",c)
	fmt.Println(s) //a
}