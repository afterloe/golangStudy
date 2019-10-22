package main

import "fmt"

type Vertex struct {
    X, Y int
}

func main(){
    v := Vertex{1, 2}
    ptr := &v
    // (*ptr).X = 1e9
    ptr.X = 1e9 // 两种方式都可以使用， 这种是语法糖
    fmt.Printf("Vertex is %v\r\n", v)
}

