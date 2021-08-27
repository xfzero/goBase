package orm_test

import (
	"fmt"

	"reflect"

	"testing"
)

type Father interface {
	animals(animal string) string

	plants(plant string) string
}

type son struct {
}

//实现接口方法

func (s *son) animals(animal string) string {

	return animal

}

func (s *son) plants(plant string) string {

	return plant

}

type sons1 struct {
	Father
}

type sons2 struct {
	Father
}

func (s1 *sons1) A(a string) {

	fmt.Println(a)

}

func (s2 *sons2) B(b string) {

	fmt.Println(b)

}

func TestOrm(t *testing.T) {

	//son才是接口的真正实现

	b := &sons1{

		Father: &sons2{Father: &son{}},
	}

	c := &sons2{

		Father: &son{},
	}

	//最终的方法都是调用son对象的方法

	fmt.Println(b.animals("两层接口继承：animals"))

	fmt.Println(c.plants("一层接口继承：plants"))

	fmt.Println(reflect.TypeOf(b.Father)) //*orm_test.sons2

	fmt.Println(reflect.TypeOf(c)) //*orm_test.sons2

	fmt.Println(b.Father.plants("一层接口继承：plants"))

	fmt.Println(b.Father.animals("一层接口继承：animals"))

	c.B("c的B方法实现") //输出结果：c的B方法实现

	//b.Father.B("Other")   //undefined (type Father has no field or method B)

}
