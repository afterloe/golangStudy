package main

import (
	"fmt"
	"errors"
)

const name = "赵小鹏" // 申明一个常量

/*
	Go语言有一些默认的行为。
		● 大写字母开头的变量是可导出的，即其他包可以读取，是公用变量;小写字母开头的不可导出，是私有变量。
		● 大写字母开头的函数也是一样，相当于class中带public关键词的公有函数;小写字母开头就是有private关键词的私有函数。
 */
func main() {
	var age = 2 // 申明一个变量
	fmt.Printf("Hello, world or %s \nage is %d \n", name, age) // 调用一个系统库 - 输出流

	var isActive = true // bool 默认值是 false
	fmt.Printf("isActive value default is %d \n", isActive)

	var a, b = 12, 24
	c := a + b
	fmt.Printf("12 + 24 = %d \n", c)

	d, e := 11.3, 23.4
	f := d * e
	fmt.Printf("11.3 * 23.4 = %.2f \n", f)

	s1 := "java"
	arr_byt_s1 := []byte(s1) // 将字符串转换为 byte 数组
	arr_byt_s1[0] = 'n' // 修改第一项为 n
	s2 := string(arr_byt_s1) // 再转换成 字符串
	fmt.Printf("修改字符串 %s\n", s2)

	s := `这是一个
	多行的文本`
	fmt.Printf("%s \n", s)

	s1 = "b" + s1[2:] // 支持字符串切割
	fmt.Printf("s1内容是 %s \r\n", s1)

	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
		fmt.Println()
	}

	/*
	Go 支持多个常量、变量、导入多个包 写法如下

	import ("fmt", "os")
	const (i=199, pi=3.1415, prefix="Go_")
	vat (i int, pi float32, prefix string)
	 */

	 const (
	 	x = iota // 0
	 	y = iota // 1
	 	z = iota // 2
		w) // 3

	fmt.Printf("%d \n", w)

	var arr [10]int // 指定数据
	arr[0] = 43
	arr[1] = 12

	fmt.Printf("the first element is %d \n", arr[0])

	arr1 := [3]int {1, 2, 3} // 带有数据的初始化
	fmt.Printf("the first element is %d \n", arr1[2])

	arr2 := [...]int {4, 5, 6} // 省略长度，自动计算
	fmt.Printf("the first element is %d \n", arr2[0])

	// 二维数组
	doubleArr := [2][4]int {{1,2,3,4}, {6,7,8,9}}
	fmt.Printf("the 1*3 的内容是 %d \n", doubleArr[1][3])

	// 动态数组 slice
	arr3 := [...]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	arr4 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	var (
		aSlice = arr3[:3]
		bSlice = arr3[4:5]
	)

	fmt.Println(arr3[3:])
	fmt.Println(arr4[4:])
	fmt.Println(arr4[2:4])

	// append函数会改变slice所引用的数组的内容，从而影响到引用同一数组的其他slice。
	fmt.Printf("aSlice first element is %b, bSlice first element is %b \n", aSlice[0], bSlice[0])

	/**
	map是无序的
	map的长度是不固定的
	 */
	//var numberMap map[string] int  // 两种申明方式
	numberMap := make(map[string]int) // make用于内建类型(map、slice和channel)的内存分配。new用于各种类型的内存分配。
	// make只能创建slice、map和channel，并且返回一个有初始值(非零)的T 类型 make返回初始化后的(非零)值。
	numberMap["one"] = 1 // 赋值
	fmt.Printf("element is %d", numberMap["one"]) // 读取数据

	numberMap = map[string]int {"c": 4, "d": 23} // 直接赋值
	_, ok := numberMap["one"] // 判断 这个 map 是否包含 one
	if ok {
		// map内置有判断是否存在key的方式，通过delete删除map的元素
		fmt.Println("one is ")
	} else {
		// map也是一种引用类型，如果两个map同时指向一个底层，那么一个改变，另一个也相应改变
		fmt.Printf("no have one . map len is %d", len(numberMap))
	}

	/*
	各类型初始值
		int		0
		int8	0
		int32	0
		int64	0
		uint	0x0
		rune	0
		byte	0x0
		float32	0
		float64	0
		boolean	false
		string	""
	 */
}
