package main

import (
	"fmt"
)

type Boy struct {
	Name string
}

func (this Boy) GetName() string {
	return this.Name
}

type Girl struct {
	Name string
}

func (this Girl) GetName() string {
	return this.Name
}

func main() {
	boy := Boy{Name:"Ap"}
	girl := Girl{Name:"Bp"}

	boyName := boy.GetName()
	girlName := girl.GetName()

	fmt.Println(boyName)
	fmt.Println(girlName)
}