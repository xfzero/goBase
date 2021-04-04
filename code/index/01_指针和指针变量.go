package main

import "fmt"

func main() {
	var name *string
	dog := "tony"
	name = &dog

	fmt.Println(dog) //tony
	fmt.Println(&dog) //0xc0000821e0
	fmt.Println(name) //0xc0000821e0

	fmt.Printf("%s\n",name)
}