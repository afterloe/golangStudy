package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main(){
    rand.Seed(time.Now().Unix()) // 类似c/c++里面的math函数一样，需要设置time才能随机
    randNum := rand.Intn(10)
    fmt.Printf("My Rand number is %d \r\n", randNum)
}
