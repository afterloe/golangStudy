package main

import "fmt"

func main(){

    var pow = []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    for i, v := range pow {
        fmt.Printf("key -> %d, value -> %v \r\n", i, v)
    }
}
