package main

import (
    "fmt"
    "io"
    "strings"
)

func main(){

    r := strings.NewReader("Hello, Reader! i'm afterloe")
    b := make([]byte, 8)
    for {
        n, err := r.Read(b)
        fmt.Printf("n= %v, err= %v b= %v \r\n", n, err, b)
        fmt.Printf("b[:n] = %q \r\n", b[:n])
        if io.EOF == err {
            break
        }
    }
}

