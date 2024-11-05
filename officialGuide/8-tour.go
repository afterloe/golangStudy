package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    Lat, Lon float64
}

// 指针操作
func (v *Vertex) Scale(f float64) {
    v.Lat = v.Lat * f;
    v.Lon = v.Lon * f;
}

// 结构体的方法
func (v Vertex) Abs() float64 {
    return math.Sqrt(v.Lat* v.Lat + v.Lon* v.Lon)
}

// 使用结构体的方法
func Abs(v *Vertex) float64 {
    return math.Sqrt(v.Lat* v.Lat + v.Lon* v.Lon)
}

func main(){
    v := Vertex{3, 4}
    fmt.Printf("string by struct -> %f \r\n", v.Abs())
    fmt.Printf("string by func -> %f \r\n", Abs(&v))
    v.Scale(10)
    fmt.Printf("string by struct -> %f \r\n", v.Abs())
}

