package main

import (
	"fmt"
	"encoding/json"
)

type Family struct {
	Relate string `json:"relate"`
	Name string `json:"name"`
}

type Person struct {
	Name string `josn:"name"`
	Age uint32 `json:"age"`
	Family []Family `json:"family"`
}

func main() {
	str := `
		{"name":"张三","age":18,"family":[{"relate":"老父亲","name":"张大"},{"relate":"大兄弟","name":"张二"}]}
	`

	var p1 Person
	if err := json.Unmarshal([]byte(str),&p1); err != nil {
		fmt.Println("Unmarshal error")
		return
	}

	fmt.Println(p1)
}