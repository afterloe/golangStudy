package main

import (
    "fmt"
    "math"
)

func compute(fn func(float64, float64) float64) float64 {
    return fn(3, 4)
}

func main(){

    // 函数亦可以当做参数进行传递
    hypot := func(x, y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }
    fmt.Printf("value is %f \r\n", hypot(5, 12))

    fmt.Printf("compute value is %f \r\n", compute(hypot))
    fmt.Printf("compute Pow value is %f \r\n", compute(math.Pow))
}

