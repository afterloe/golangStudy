package main

import  (
	"./fs" // 导入函数 会调用 该包的 init 函数
	"fmt"
	"strconv" // 工具类 转换 bool 转换为 string
	"reflect"
)

func main() {
	fs.ReadFile("../info") // 调用方法
	//fs.sayHello()  go 约定 大写开头的就是 public 小写开头的 就是 private
	fmt.Println("export value is " + strconv.Itoa(fs.MIN_LINE)) // 获取 fs 导出的变量

	// 调用包里的模块
	info, err := fs.ReadRealFile("/Users/afterloe/Afterloe/go/golangStudy", ".gitignore")
	// 如果存在error 就输出
	if nil != err {
		fmt.Println("find err")
		fmt.Println(err.Error()) // 输出error
		return
	}

	fmt.Println(info)

	flag, err := fs.WriteRealFile([]byte("i need you"), "test")
	if err != nil {
		fmt.Println("find error")
		fmt.Println(err.Error())
	}

	fmt.Println("write file " + strconv.FormatBool(flag))

	user := map[string]interface{}{"name": "afterloe", "sex": 0, "domain": [...]string{"go", "java", "cpp"}}
	userStr, error := fs.FormatToString(user)
	if nil != error {
		fmt.Println(error)
	}
	fmt.Println(userStr)

	json, _ := fs.ReadRealFile("/Users/afterloe/Afterloe/node/cynomy_axure", "package.json")
	obj,_ := fs.FormatToStruct(json)

	depMap := obj["dependencies"]
	v := reflect.ValueOf(depMap)
	// index
	for _, key := range v.MapKeys() {
		val := v.MapIndex(key)
		fmt.Printf("%s -> %s \n", key.Interface(), val.Interface())
	}
}