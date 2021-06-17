package main

import "fmt"

func main() {
	c := 'a'
	fmt.Println(c) //97 字符类型是特殊的整形
	fmt.Printf("%c\n",c) //a

	s := fmt.Sprintf("%c",c)
	fmt.Println(s) //a

	arr := [...]int{1,2,4,5}
	//arr := []int{1,2,4,5} //slice
	for _,v := range arr {
		fmt.Println(v)
	}
	fmt.Printf("%T\n",arr)

	arr2 := [3][2]int{
		{1,2},
		{2,3},
		{3,4}, //逗号必须加
	}
	for _,v := range arr2 {
		fmt.Println(v)
	}
	
}
