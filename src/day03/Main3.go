package main

import "fmt"

// 参数接收的时候也是一样
func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	// 通过操作符<-来接收和发送数据。
	c <- sum // send sum to c
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// chan 发送或者接收值。这些值只能是特定的类型:channel类型。
	// 定义一个channel时，也需要定义发送到channel的值的类型。
	// 必须使用make创建channel。
	c := make(chan int)
	go sum(a[:len(a)/2], c) // 数组后一半
	go sum(a[len(a)/2:], c) // 数组前一半 ':' 截断操作
	// channel接收和发送数据都是阻塞的 无缓冲channel是在多个goroutine之间同步很棒的工具
	x, y := <- c, <- c // return from c 通过操作符<-来接收和发送数据。
	fmt.Println(x, y, x + y)

	// 申明一个buffered Channels
	c2 := make(chan int, 3)  // value == 0 无缓冲 柱塞，value > 0 缓冲，直到填满value
	// 这个channel中，前3个元素可以无阻塞的写入。
	// 当写入第4个 元素时，代码将会阻塞，直到其他goroutine从channel中读取一些元素才能放入
	c2 <- 1
	fmt.Println(<- c2)
	c2 <- 2
	c2 <- 3
	c2 <- 4
	fmt.Println(<- c2)
	fmt.Println(<- c2)
	fmt.Println(<- c2)
	//fmt.Println(<- c2) // 超出返回是会报错的 all goroutines are asleep - deadlock!
}
