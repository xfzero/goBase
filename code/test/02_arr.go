package main

import "fmt"

func main() {
	// var arr [4]int
	// arr := [4]int{1,2,4,5}
	// arr := [4]int{1} //其他元素为0
	arr := [...]int{1,2,4,5}
	// arr := []int{1,2,4,5} //slice
	for _,v := range arr {
		fmt.Println(v)
	}
	fmt.Printf("%T\n",arr)

	//二维数组
	/*var b [3][2]string
    b[0][0] = "apple"
    b[0][1] = "samsung"
    b[1][0] = "microsoft"
    b[1][1] = "google"
    b[2][0] = "AT&T"
    b[2][1] = "T-Mobile"
    fmt.Printf("\n")*/
	arr2 := [3][2]int{
		{1,2},
		{2,3},
		{3,4}, //逗号必须加
	}
	for _,v := range arr2 {
		fmt.Println(v)
	}
	
}