package main

import (
    "fmt"
    "sync"
    "time"
)

// Go协程能够访问一个共享的变量
type SafeCounter struct {
    v    map[string]int
    mux  sync.Mutex //  sync.Mutex 互斥锁 Lock、Unlock
}

func (c *SafeCounter) Inc(key string) {
    c.mux.Lock()
    c.v[key]++
    c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
    c.mux.Lock()
    defer c.mux.Unlock()  // 解锁
    return c.v[key]
}

func main() {
    c := SafeCounter{v: make(map[string]int)}
    for i := 0; i < 1000; i++ {
        go c.Inc("somekey")
    }

    time.Sleep(time.Second)
    fmt.Println(c.Value("somekey"))
}
