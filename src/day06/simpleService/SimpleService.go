package simpleService

import (
	"net/http"
	"fmt"
	"strings"
)

// 设置允许访问的域
const POINT = "127.0.0.1:8082"

// 路由函数
func Execute(response http.ResponseWriter, request *http.Request) {
	// 格式化输入参数 - 默认是不会格式化的
	request.ParseForm()
	// 输出一些内容
	fmt.Println(request.Form) // 返回是 map[string][]interface{}
	fmt.Println(request.URL.Path) // 返回的是相对路径 /v1/user
	fmt.Println(request.URL.Scheme)
	// 这TM 返回的是一个 []interface{}
	fmt.Println(request.Form["name"])

	for key, val := range request.Form {
		fmt.Printf("%s -> %s \n", key, strings.Join(val, " | "))
	}
	// 向输出流写入输出内容
	fmt.Fprintf(response, "hello %s", request.Form["name"])
}


