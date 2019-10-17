package main

import (
	"fmt"
)

type Animal struct {
	name string
	age int
}

type Dog struct {
	Animal
}

type Cat struct {
	Animal
}

// 使用 占位符 但是不能调用原来的内容
func (_ Animal) say() {
	fmt.Println("吼~~")
}

// 给struct 绑定方法
func (d Dog) say() {
	// 调用d的参数
	fmt.Println(d.name, "say 汪汪汪")
}

// 可以使值传递 也可以是 指针传递
// 指针会对实例对象的内容发生操作，而普通类型是以副本作为操作对象，并不对原实例对象发生操作
func (c *Cat) say() {
	fmt.Println(c.name,"say 喵喵喵")
}

// method是可以再任何自定义的类型、内置类型、struct等各种类型上的
// 其他自定义类型申明如下
// type typeName typeLiteral
type months map[string]int

func main() {
	littleDog := Dog{Animal{"小黑", 2}}
	littleDog.say()
	littleCat := Cat{Animal{"小花", 1}}
	littleCat.say()
	littleDog = Dog{}
	littleDog.Animal.say()

	// 使用自定义类型
	m := months{"January": 31, "February": 28}
	for key, value := range m {
		fmt.Println(key, " --> ", value)
	}
}
