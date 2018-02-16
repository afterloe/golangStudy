package simpleService2

import (
	"fmt"
	"net/http"
)

type CustomHandler struct {
	ServerListenAddr string
}

func Execute(response http.ResponseWriter, request *http.Request) {
	// 格式化输入参数 - 默认是不会格式化的
	request.ParseForm()
	// 输出一些内容
	fmt.Println(request.Form)
	fmt.Println(request.Method) // 请求方式

	if 0 == len(request.Form["name"]) {
		fmt.Fprintln(response, "lack parameter -> name")
		return
	}

	// 向输出流写入输出内容
	fmt.Fprintf(response, "hello %s", request.Form["name"][0])
}

func (*CustomHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// 只处理 某个路径
	if request.URL.Path == "/" {
		Execute(response, request)
		return
	}

	http.NotFound(response, request)
	return
}