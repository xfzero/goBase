package main

import "fmt"

func main() {
	//切片都有一个底层对应的数组，修改切片的元素实际修改的就是对应数组的值
	var a []int
	fmt.Println(a)

	b := []int{6,7,8}
	fmt.Println(b)

	c := [5]int{1,2,3,4,5}
	d := c[0:3]
	fmt.Println(d) //[1,2,3]
	fmt.Println(d[0])
	d = append(d,6)
	fmt.Println(c) //[1,2,3,6,5]
	d = append(d,7)
	fmt.Println(c) //[1,2,3,6,7]
	c[1] = 8
	fmt.Println(c) //[1,8,3,6,7]

	e := make([]int,len(d))
	copy(e,d)
	e[1] = 9
	fmt.Println(c) //[1,8,3,6,7] 和php中的对象类似，copy后e重新管理一个新数组 所以修改e后c数组不变，如果是赋值则会变(e=d)
	d[1] = 9
	fmt.Println(c) //[1,9,3,6,7]


	f := []int{1,2,3,4,5}
	index := 2
	f = append(f[:index],f[index+1:]...)
	fmt.Println(f) //[1,2,4,5]
	fmt.Println(f[2]) //4
}