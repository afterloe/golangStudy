package main

import (
	"fmt"
	"time"
	"runtime"
)

func fibonacci (n int, c chan int) {
	x, y:= 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}
	close(c) //  显式的 关闭chan
	// 在生产者的地方关闭channel
}

func seance(c, quit chan int) {
	x, y := 1, 1
	// select不再阻塞等待channel 所以需要死循环
	for {
		// select其实就是类似switch的功能
		select {
		case c <- x:
			x, y = y, x + y
		// 有时候goroutime会有阻塞的时候，可以使用select来设置超时操作
		case <- time.After(5 * time.Second):
			fmt.Println("time out")
			break
		case <- quit:
			fmt.Println("quit")
			return
			// 在select里面还有default语法，default就是当监听的channel都没有准备好的时候，默认执行的
		}
	}
}

func main()  {
	c := make(chan int, 12) // 需要指定长度
	go fibonacci(cap(c), c)
	if _, ok := <- c; ok { //
		fmt.Println("channel 没有关闭")
	}
	// rang c 能够不断地读取channel里的数据，直到该channel被显示的关闭
	for i := range c {
		fmt.Println(i)
	}

	// 当 channel 关闭的时候是不能在向其发送数据的
	b, quit := make(chan int), make(chan int)
	// 异步执行匿名函数
	go func() {
		for i := 0; i< 10; i++ {
			fmt.Println(<-b)
		}
		quit <- 0
		runtime.Goexit() // 退出当前执行的goroutine
	}()
	seance(b, quit)

	fmt.Println(runtime.NumCPU()) // 输出CPU核数
	fmt.Println(runtime.NumGoroutine()) // 输出正在执行和排队的任务总数
}
