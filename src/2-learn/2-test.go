package main

import (
	core "fmt" // 这种导入是可以指定 导入包的名字的，如果不指定。默认是用 包名
	//_ "fmt" // 这种不是导入，而是调用了 fmt 里面的 init 函数
)

/*
定义 struct
 */
type person struct {
	name string
	mail string
	age int8
}

func personToString (p *person) {
	core.Println(p.name)
	core.Println(p.mail)
	core.Println(p.age)
}

func olderMan(p1, p2 *person) (older *person) {
	if p1.age > p2.age {
		older = p1
	} else {
		older = p2
	}
	return
}

/*
	go 是支持只提供类型，二不写字段的方式来申明struct
 */
type developer struct {
	person	// 可以理解为继承吧。。。 包含了person的所有字段
	domain string
	//age int8
}

func main() {
	core.Println("指定别名 导入外部包")
	// struct 使用 - 第一种使用方式
	var me person
	me.name = "afterloe"
	me.mail = "lm6289511@gmail.com"
	me.age = 24

	// 第二种使用方式，这种必须要补全 - 否则会报错
	he := person{"afterloe.L","lm6289511@gmail.com", 21}
	personToString(&he) // 打印出person的内容

	// 第三种，可以缺省
	he2 := person{mail: "lm6289511@gmail.com", age: 17}
	older := olderMan(&he, &he2) // 注意这里返回的是 person 的指针
	personToString(older)


	older.age = 4 // 指针修改的是 执行的内容，如果发生修改 所有的东西都进行了修改的
	personToString(&he)

	goLangDeveloper := developer{he2, "golang"} // 申明一个程序员
	personToString(&goLangDeveloper.person)
	core.Println("worker age is ", goLangDeveloper.age) // 字段是可以直接访问的！
	// 如果覆盖的话 ，以最外层的为准。 当然如果想访问最原来的值可以用下面的方法
	core.Println("golang developer age is ", goLangDeveloper.person.age)
	core.Println("developer domain is ", goLangDeveloper.domain)
}
