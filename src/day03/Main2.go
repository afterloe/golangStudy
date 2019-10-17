package main

import (
	"fmt"
	"runtime"
)

func say(word string) {
	for i:=0; i < 10; i++ {
		runtime.Gosched() // 让CPU把时间片让给别人，下次某个时候继续恢复执行该goroutine
		fmt.Println(word)
	}
}

func main() {
	/*
	默认情况下，调度器仅使用单线程，也就是说只实现了并发。
	想要发挥多核处理器的并行，
	需要在我们的程序中显示调用 runtime.GOMAXPROCS(n)告诉调度器同时使用多个线程。
	GOMAXPROCS设置了同时运行逻辑代码的系统线程的最大数量，并返回之前的设置。
	如果n < 1，不会改变当前设置
	 */
	 // 都次运行 返回的结果都不一样
	go say("afterloe") // 异步执行
	say("joe") // 当前 执行
}
