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


