1. 指针和指针变量

```go
var name *string
var dog := "tony"
name = &dog
fmt.Println(dog) //tony
fmt.Println(&dog) //0xc0000821e0
fmt.Println(name) //0xc0000821e0
```

new(T):创建一个变量T，初始值为零值，返回变量的指针
```go
p := new(int)
fmt.Println(*p) //0 
```


2. type关键字

3. interface{}返回值类型
golang对于不确定返回值可以用interface{}代替，这确实很方便，但是也带来了问题，那就是如何判断返回值是什么类型的？其实可以用反射也就是reflect来判断

```go
func HandleRequest() interface{} {
	arr := [...]int{1,2,4,5}
	return arr
}
```

4. 卫述语句
a, b := 4, 0
res, err := divisionInt(a, b)
if err != nil {
   fmt.Println(err.Error())
   return
}
fmt.Println(a, "除以", b, "的结果是 ", res)

5. 空白标识符(_)

6. array,slice,map

7. p,&p,*p

8. mod & dep

9. Makefile

10. struct
golang非面向对象语言，可以使用struct实现弱面向对象的特性
继承、重载、多态

new对象：
``go
   //user.go
   type User struct{
      Name string
   }
   func newUser(id uint32, name string) *User {
      return &User{
         id: id,
         name: name,
      }
   }
``

//继承
type Animal strcut {
   Color string
}
type Duck strcut {
   Animal
}

11. defer,panic,recover
三者一起使用可以充当其他语言中try…catch…的角色，而defer本身又像其他语言的析构函数

12. rune,byte,string

13. go run -race main.go
