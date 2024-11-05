package main

import "fmt"

func add(x, y, sum *int) {
    fmt.Printf("function add x -> %d \r\n", *x)
    fmt.Printf("function add y -> %d \r\n", *y)
    *sum = *x + *y
}

func main() {
    var (
        x = 10
        y = 22
    )
    sum := new(int)
    add(&x, &y, sum)
    fmt.Printf("10 + 22 = %d \r\n", *sum)
}
