package main

import (
	"fmt"
)

//类型断言

type People interface {
	ReturnName() string
	ReturnAge() uint32
}

type Student struct {
	Name string
	Age uint32
}

func (this Student) ReturnName() string {
	return this.Name
}

func (this Student) ReturnAge() uint32 {
	return this.Age
}

//类型断言
func CheckPeople(test interface{}) {
	if _, ok := test.(People); ok {
		fmt.Println("Student implements People")
	} else {
		fmt.Println("Student not implements People")
	}
}

func main() {
	cbs := Student{Name:"tom",Age:18}

	CheckPeople(cbs) // Student implements People


	Params := make([]interface{}, 3)
	Params[0] = 88                   // 整型
	Params[1] = "咖啡色的羊驼"         // 字符串
	Params[2] = Student{Name: "cbs"} // 自定义结构体类型
	
	// Comma-ok断言
	for index, v := range Params {
		if _, ok := v.(int); ok {
			fmt.Printf("Params[%d] 是int类型 \n", index)
		} else if _, ok := v.(string); ok {
			fmt.Printf("Params[%d] 是字符串类型\n", index)
		} else if _, ok := v.(Student); ok {
			fmt.Printf("Params[%d] 是自定义结构体类型\n", index)
		} else {
			fmt.Printf("list[%d] 未知类型\n", index)
		}
	}
	
	// switch判断
	for index, v := range Params {
		switch  value := v.(type) {
        case int:
            fmt.Printf("Params[%d] 是int类型, 值：%d \n", index,value)
        case string:
            fmt.Printf("Params[%d] 是字符串类型, 值：%s\n", index,value)
        case Student:
            fmt.Printf("Params[%d] 是Person类型, 值：%s\n", index,value)
        default:
            fmt.Printf("list[%d] 未知类型\n", index)
        } 
	
	}  
}
