1. 指针和指针变量

   ```go
   var name *string
   var dog := "tony"
   name = &dog
   fmt.Println(dog) //tony
   fmt.Println(&dog) //0xc0000821e0
   fmt.Println(name) //0xc0000821e0
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


