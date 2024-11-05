package main

import (
    "fmt"
    "math"
)

type Abser interface {
    Abs() float64
}

func main() {
    var ptr Abser
    f := MyFloat(-math.Sqrt2)
    v := Vertex{3, 4}

    ptr = f
    fmt.Printf("myfloat value is %f \r\n", ptr.Abs())
    // 因为Vertex的指针实现了 Abs的方法，所以直接使用方法是无法调用的。
    ptr = &v
    fmt.Printf("myfloat value is %f \r\n", ptr.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
    if 0 > f {
        return float64(-f)
    } else {
        return float64(f)
    }
}

type Vertex struct {
    X, Y float64
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
