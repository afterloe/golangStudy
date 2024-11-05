package main

import "fmt"

func do(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("this is int, Twice %v is %v \r\n", v, v*2)
    case string:
        fmt.Printf("this is string, %q is %v bytes long \r\n", v, len(v))
    default:
        fmt.Printf("I don't know about this type %T! \r\n", v)
    }
}

func main(){
    do(21)
    do("hello")
    do(true)
}

