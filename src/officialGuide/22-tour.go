package main

import "fmt"

func sum(s []int, c chan int) {
    sum := 0
    for _,v := range s {
        sum += v
    }
    c <-sum
}

func main(){
    s := []int{7, 2, 8, -9, 4, 2}

    c := make(chan int)
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)
    x, y := <-c, <-c

    fmt.Printf("%d\t%d\t%d\r\n", x, y, x+y)
}

