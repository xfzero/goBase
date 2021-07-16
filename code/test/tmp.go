package main

import (
	"fmt"
	"os"
)

func main() {
	aa := 12;bb := 13;
	fmt.Println(aa,bb)
	filename := "./gm_config.xml"
	file, err := os.Open(filename)
	if err != nil {
        fmt.Println("Open file Failed", err)
        return
    }
    defer func() {
        file.Close()
        fmt.Println("file close1")
    }()
    //return

    var b []byte = make([]byte, 4096)
    n, err := file.Read(b)
    if err != nil {
        fmt.Println("Open file Failed", err)
    }
    data := string(b[:n])

    fmt.Println("222222")
    defer func() {
        file.Close()
        fmt.Println("file close2")
    }()
    return
    fmt.Println(data)
}

