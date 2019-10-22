package main

import "fmt"

type Vertex struct {
    Lat, Lon float64

}

func main(){

    var m map[string]Vertex
    m = make(map[string]Vertex)
    m["福州"] = Vertex{26.08, 119.28}
    m["泉州"] = Vertex{24.88, 118.67}
    m["深圳"] = Vertex{22.32, 114.03}
    fmt.Printf("m -> %v \r\n", m)
    fmt.Printf("福州 -> %v \r\n", m["福州"])

    // 第二中写法
    var m1 = map[string]Vertex{
        "厦门": Vertex{24.48, 118.08},
        "兰溪": {29.19, 119.48}, // 这样写也是可以的
        "漳州": Vertex{23.82, 117.12},
    }
    fmt.Printf("m1 -> %v \r\n", m1)
    v, ok := m1["厦门"]
    fmt.Printf("value -> %v, flag is %v \r\n", v, ok)
    delete(m1, "厦门") // 删除key
    v, ok = m1["厦门"]
    fmt.Printf("value -> %v, flag is %v \r\n", v, ok)
}

