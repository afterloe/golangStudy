package simpleService2

import (
	"fmt"
	"net/http"
)

type CustomHandler struct {
	ServerListenAddr string
}

func Execute(response http.ResponseWriter, request *http.Request) {
	// 格式化输入参数 - 默认是不会格式化的 这个格式化只需要执行一次就可以了。
	// 在使用前格式化，多次调用时没有效果的 eg: if r.Form == nil
	request.ParseForm()
	// 输出一些内容
	fmt.Println(request.Form)
	fmt.Println(request.Method) // 请求方式

	// 自己写一个 函数Scheme 的校验规则
	if 0 == len(request.Form["name"]) {
		fmt.Fprintln(response, "lack parameter -> name")
		return
	}

	// 向输出流写入输出内容
	fmt.Fprintf(response, "hello %s", request.Form["name"][0])
}

func (*CustomHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// 只处理 某个路径 这里可用正则来扩充
	if request.URL.Path == "/" {
		Execute(response, request)
		return
	}

	http.NotFound(response, request) // 自己也是可以编写 404 这类的返回的 源码只是 func NotFound(w ResponseWriter, r *Request)
	return
}