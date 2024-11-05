package main

import "fmt"

func main(){
    var i interface{} = "afterloe"

    s := i.(string)
    fmt.Println(s)

    // 推荐的 interface 类型转换
    s, ok := i.(string)
    fmt.Println(s, ok)

    f, ok := i.(float64)
    fmt.Println(f, ok)

    // 会报错 panic
    f = i.(float64)
    fmt.Println(f)
}

