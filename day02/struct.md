# 结构体 struct

> 代码链接 [./Main2.go](./Main2.go)

struct不仅能将struct作为匿名字段，自定义类型、内置类型都可以作为匿名字段，而且可以在相应的字段上进行函数操作

```golang
/*
	go 是支持只提供类型，二不写字段的方式来申明struct
 */
type developer struct {
	person	// 可以理解为继承吧。。。 包含了person的所有字段
	domain string
	//age int8
}

	goLangDeveloper := developer{he2, "golang"} // 申明一个程序员
	personToString(&goLangDeveloper.person)
	core.Println("worker age is ", goLangDeveloper.age) // 字段是可以直接访问的！
	// 如果覆盖的话 ，以最外层的为准。 当然如果想访问最原来的值可以用下面的方法
	core.Println("golang developer age is ", goLangDeveloper.person.age)
	core.Println("developer domain is ", goLangDeveloper.domain)
```

具体可以参考源代码的内容
