package main

import (
	"fmt"
	"reflect"
)

type voidI interface {}

func getType(i voidI) {
	t := reflect.TypeOf(i) // 获取空函数的类型
	fmt.Println(t) // 输出
	v := reflect.ValueOf(i) // 获取值
	fmt.Println(v)
}

/*
interface 的字段是不可以修改的

func change(i *voidI) {
	p := reflect.ValueOf(i) // 获取值的地址
	v := p.Elem() // 转换为  Elem对象
	v.SetFloat(7.5) //
}
*/

func main() {
	var x = 4.5
	getType(x) // float64
	var d = "afterloe"
	getType(d) // string

	// 如果类型不一致是会报错的
	//change(x)
	//fmt.Println(x)

	//var i voidI
	//i = 3.5
	//change(&i)
	fmt.Println(x)
	p := reflect.ValueOf(&x)
	v := p.Elem()
	v.SetFloat(3.14)
	//v.SetString("afterloe") // 类型必须是一致的才可以修改
	getType(x)
}
