package main

import (
	"fmt"
)

type Test struct {
	name string
}

type vo1 interface {
	equal()
}

// 嵌入式 interface 理解为 interface 的继承
type vo2 interface {
	vo1
	say(vo1) string // 也可以使用 函数传参的形式
}

func (t Test) equal() {
	fmt.Println("equal " + t.name)
}

func (t Test) say(v vo1) string{
	fmt.Println("---------------------------")
	v.equal()
	fmt.Println("say ...")
	return t.name
}

func main() {
	t := Test{name: "afterloe"}
	var v1 vo1
	v1 = t
	v1.equal()

	t1 := Test{name: "Joy"}
	var v2 vo2
	v2 = t1
	v2.say(v1) // 嵌入式函数调用
	v2.equal()
}