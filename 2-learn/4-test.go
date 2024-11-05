package main

import (
	"fmt"
	"strconv" // Itoa 方法可以把 int 转换为 string
)

type Person struct {
	name string
	age int
}

type Developer struct {
	Person
	domain string
	tags map[string] int
}

// 申明一个 interface 包含两个方法
type HumanActivity interface {
	say(word, time string)
	doSomething()
}

// 申明一个空的interface 他是可以存储接收任意数值，可以理解为 void *
type voidActivity interface {

}

func (p Person) say (word, time string) {
	fmt.Printf("%s\t[%s]: %s \n", p.name, time, word)
}

// 自己实现
func (_ Person) doSomething() {
	fmt.Println("want to do something ...")
}

func (developer Developer) say(word, time string) {
	fmt.Printf("程序员[%s][%s]: %s \n", developer.domain, time, word)
}

/**
重写 toString 方法
 */
func (p Person) String() string {
	return "name is " + p.name + " age is " + strconv.Itoa(p.age)
}

func (developer Developer) doSomething() {
	fmt.Printf("%s -------------------------- \n", developer.domain)
	for key, val := range developer.tags {
		fmt.Printf("%s -> %d \n", key, val)
	}
}

func main() {
	joe := Person{name: "joe", age: 6}
	yyy := Person{name: "yyy", age: 5}
	grace := Developer{Person{name: "grace", age: 25}, "java", map[string]int {"j2ee":2, "j2me":5}}
	tom := Developer{Person{name: "grace", age: 25}, "go", map[string]int {"info": 0}}
	time := "2018-01-30 14:12:15"

	// 申明一个interface 变量
	var activity HumanActivity
	activity = joe  // 如果这个struct 有这个方法(参数列表也要一致 就能赋值)
	activity.say("今天", time)
	activity.doSomething()  // 直接调用方法

	activity = grace
	activity.say("明天", time)
	activity.doSomething()

	actList := make([]HumanActivity, 3)
	actList[0], actList[1], actList[2] = yyy, tom, grace // 也能用 slice 进行结构

	// 支持循环调用
	for _, value := range actList {
		value.doSomething()
	}

	// 空 interface 的运用
	i := 10
	var va voidActivity
	va = i
	s1 := "hello world"
	va = s1
	fmt.Printf("%s \n", va)
	fmt.Println("joe is ", joe) // 调用默认的toString 方法
	fmt.Println("tom is ", tom.String()) // 调用父级的toString 方法

	// comma-ok 断言
	// 是否确实 存储了T类型的数值
	// value就是变量的值，ok是一个bool类 型，element是interface变量，T是断言的类型
	if value, ok := activity.(Developer); ok {
		fmt.Println("activity 是Developer 值是", value)
	}

	// 如果element里面确实存储了T类型的数值，那么ok返回true，否则返回false。
	for _, element := range actList {
		// 判断 数组中的 interface 是不是特定的类型
		if value, ok := element.(Developer); ok {
			fmt.Println("has info", value, ok)
		}
	}
	/*
	// 还可以用switch case 来处理 不过 element.(type)语法不能在switch外使用
	switch value := element.(type) {
		case Developer:
		break;
		case Person;
		break;
		default:
		break;
	}
	 */
}
