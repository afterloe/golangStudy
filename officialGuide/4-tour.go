package main

import "fmt"

type Vertex struct {
    X,Y int
}

func main(){
    // 结构体的申明方式
    var (
        v1 = Vertex{11, 23}
        v2 = Vertex{X: 12}
        v3 = Vertex{}
        v4 = &Vertex{22, 24}
    )
    fmt.Printf("%v\r\n%v\r\n%v\r\n%v\r\n", v1, v2, v3, v4)
}

