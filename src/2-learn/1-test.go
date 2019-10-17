package main

import "fmt"

/**
	函数 定义如下

	func funcName(input1 type, input2 type) (output1 type, output2 type) {
		// ... do something
		return value1, value2
	}

	go里面可以为函数的返回值申明变量，当然这种写法也是可以的

	func funcName(input1 type, input2 type) (type, type) {
	如果没有返回值 可以忽略 return、有的话 必须写上 return
 */
func sum(a, b int) int {
    return a + b
}

 /**
 	多参数返回, 简写
  */
func sumAndMax(a, b int) (int, int) {
    max := 0
    if a > b {
        max = a
    } else {
        max = b
    }

    return sum(a, b), max
}
 /**
 	官方建议
func sumAndMax2(a, b int)(sum int, max int)  {
	// do something
	return
}
*/

/* 可变参数
func myFunc(arg ... int) {
	for _, n : range arg {
		fmt.Printf("And the number is: %d \n", n)
	}
}
 */
/* 值传递和指针传递 别忘了
func add1(a *int) int {
	// 这是会修改指针指向的值，基本就是全局修改了。而值传递是不会进行修改的。只是copy一份数据
}

传指针使得多个函数能操作同一个对象。
传指针比较轻量级(8bytes)，只是传内存地址
string，slice，map这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针
 */

 /*
 defer 延迟(defer)语句，你可以在函数中添加多个defer语句。当函数执行到最后时，这些defer语句会 按照逆序执行，最后该函数返回
 进行一些打开资源的操作时，遇到错误需要提前返回，在返回前你需要关闭相应的 资源，不然很容易造成资源泄露等问题

 func readFile() bool {
 	file.Open("file")
 	if errX {
 		file.close()
 		return false;
 	}
 	if errY {
 		file.close()
 		return false;
 	}
 	file.close()
 	return true;
 }
 // 优化后
 func readFile() bool {
 	file.Open("file")
 	defer file.Close() // defer 会在函数退出前调用
 	if errX {
 		return false;
 	}
 	if errY {
 		return false;
 	}
 	return true;
 }
 // 如果有很多调用defer，那么defer是采用后进先出模式

打开文件
做一些事情
关闭文件
记录日志
 */

func mockReadFile() {
	fmt.Println("打开文件")
	defer fmt.Println("记录日志")
	defer fmt.Println("关闭文件")
	fmt.Println("做一些事情")
}

// 同脚本语言一样，golang 也是可以将函数 作为值、类型进行传递的 通过type来定义它，它的类型就是所有拥有相同的参数，相同的返回值
/*
type typeName func(input1 type, input2 type) (return1 type, return2 type)
 */
type testInt func (int, int) int // 定义一个函数类型

// 函数当做值和类型 写一些通用接口的时候非常有用
func useType(a int, b int, f testInt) (value int) {
	value = f(a, b)
	return
}

// go 异常处理
/*

 */

func main() {
	var (
		a = 13
		b = 11
	)

	fmt.Println("a > b ? => ", a > b)

	// 计算获取c的值 然后判断c的值是不是大于10
	if c := a + b; c > 10 {
		fmt.Println("a > b", c)
		// goto 语法 标签名是大小写敏感的
		goto GOThere
	} else {
		fmt.Println("a < b", c)
	}

	GOThere: // goto 标签
		fmt.Printf("sum is %d \n", sum(a, b))
		fmt.Println("i go here")
		//fmt.Println(c) 外部是不能使用这个c的 他的域不在这. 但是这里是可以的，因为这个c不是上面那个
		c , d := sumAndMax(a, b)
		fmt.Printf("sum is %d, max is %d \n", c, d)

	// 简单的循环逻辑
	sumValue := 0
	for sumValue < 100 {
		sumValue+=1
	}
	fmt.Printf("sum is %d \n", c)

	// for 也可以配合 range 用于读取 splice 和 map的 数据
	simpleMap := map[string]int { "one":1, "two":2, "three":3 }
	for key, value:= range simpleMap {
		fmt.Printf("simpleMap key is %s \n", key)
		fmt.Printf("simapleMap Value is %d \n", value)
	}

	// switch 用法大多地方是一致的
	switch a {
	case 10:
		fmt.Println("this is 10")
	default:
		fmt.Println("no value is show.")
	}

	mockReadFile()
	fmt.Printf("value is %d", useType(1001, 20031, sum))
}
