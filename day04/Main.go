package main

import  (
	"./fs" // 导入函数 会调用 该包的 init 函数
	"fmt"
	"strconv"
)

func main()  {
	fs.ReadFile("../info") // 调用方法
	//fs.sayHello()  go 约定 大写开头的就是 public 小写开头的 就是 private
	fmt.Println("export value is " + strconv.Itoa(fs.MIN_LINE)) // 获取 fs 导出的变量

	// 调用包里的模块
	info, err := fs.ReadRealFile("/Users/afterloe/Afterloe/go/golangStudy", "LICENSE")
	// 如果存在error 就输出
	if nil != err {
		fmt.Println("find err")
		fmt.Println(err.Error()) // 输出error
		return
	}

	fmt.Println(info)
}
