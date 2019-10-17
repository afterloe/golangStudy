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

	flag, err := fs.WriteRealFile([]byte("i need you"), "test") // 往 /tmp/test 文件中写入一句话
	if err != nil {
		fmt.Println("find error")
		fmt.Println(err.Error())
	}

	fmt.Println("write file " + strconv.FormatBool(flag))

	user := map[string]interface{}{"name": "afterloe", "sex": 0, "domain": [...]string{"go", "java", "cpp"}}
	userStr, error := fs.FormatToString(&user) // 获取user转换后的 json文本
	if nil != error {
		fmt.Println(error)
	}
	fmt.Println(userStr)

	// 读取本地文件中的json文件
	json, _ := fs.ReadRealFile("/Users/afterloe/Afterloe/node/cynomy_axure", "package.json")
	obj, _ := fs.FormatToStruct(&json) // 文件转换为 vol interface {} 的实例

	//depMap := obj["dependencies"].(map[string]interface {}) // 可以直接使用 获取的也是 vol interface {}
	depMap := obj["dependencies"] // 可以直接使用 获取的也是 vol interface {}
	fmt.Println(obj["name"].(string)) // 按照string 进行处理
	fmt.Println(obj["age"].(float64)) // 转换为int 进行处理
	v := reflect.ValueOf(depMap) // 使用反射 获取 map

	// 获取map中的指定值
	// fmt.Println(depMap["koa"])
	fmt.Println(reflect.TypeOf(depMap)) // map[string]interface {}
	fmt.Println(v.MapIndex(reflect.ValueOf("koa"))) // 使用这个方式来直接取map 中 koa的值

	// index 第一个参数是 下标 -- 循环输出map中的值
	for _, key := range v.MapKeys() {
		val := v.MapIndex(key)
		fmt.Printf("%s -> %s \n", key.Interface(), val.Interface())
	}

	// 循环输出 slice 中的值
	keywords := obj["keywords"].([]interface{})
	fmt.Println(reflect.TypeOf(keywords))
	v = reflect.ValueOf(keywords)
	fmt.Println(v.Index(0)) // 直接取第一个数的值
	for _, val := range keywords{
		val := reflect.ValueOf(val)
		fmt.Println(val.Interface())
	}
}