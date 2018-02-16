package simpleService2

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"time"
)

type CustomHandler struct {
	ServerListenAddr string
}

func UploadPic(response http.ResponseWriter, request *http.Request) {
	if "POST" != request.Method {
		response.Write([]byte("Method is not supper! Please use POST"))
		return
	}

	// 设置转换文件的内存大小 - 即 buffer 的大小
	request.ParseMultipartForm(5 << 20) // 5MB 单位是 bytes  同理 32 << 10 32kb
	file, handler, error := request.FormFile("icon")
	if error != nil {
		response.Write([]byte("save file find error !" + error.Error()))
		return
	}

	defer file.Close()
	fmt.Println(handler.Header) // 输出文件的请求头
	tf, error := os.OpenFile("/Users/afterloe/Documents/" + handler.Filename, os.O_WRONLY | os.O_CREATE, 0666)
	if error != nil {
		response.Write([]byte("save file find error !" + error.Error()))
		return
	}
	defer tf.Close()
	io.Copy(tf, file)
	response.Write([]byte("upload success -> " + time.Now().String()))
}

func Execute(response http.ResponseWriter, request *http.Request) {
	// 格式化输入参数 - 默认是不会格式化的 这个格式化只需要执行一次就可以了。
	// 在使用前格式化，多次调用时没有效果的 eg: if r.Form == nil
	request.ParseForm()
	// 输出一些内容
	fmt.Println(request.Form)
	fmt.Println(request.Method) // 请求方式

	/*
	// 一种用于多选循环的库 一般用于多选， select等表单的时候进行判断
	slice := []string{"afterloe", "joe"}
	for _, val := rang slice {
		if val == request.Form.Get("name") {
			return true
		}
	}
	 */
	fmt.Println(request.Form.Get("name")) // 也可以使用这种方式来选择

	// 自己写一个 函数Scheme 的校验规则
	if 0 == len(request.Form["name"]) {
		fmt.Fprintln(response, "lack parameter -> name")
		return
	}

	// 向输出流写入输出内容
	fmt.Fprintf(response, "hello %s", request.Form["name"][0])
}

/*
html/template有以下几个函数可以转义。用于防止html注入
	● func HTMLEscape(w io.Writer, b []byte) //把b进行转义之后写到w
	● func HTMLEscapeString(s string) string //转义s之后返回结果字符串
	● func HTMLEscaper(args ...interface{}) string //支持多个参数一起转义，返回结果字符串
 */
func (*CustomHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// 只处理 某个路径 这里可用正则来扩充
	// regexp.MatchString("", request.URL.Path) 正则来匹配
	path := request.URL.Path
	if path == "/" {
		Execute(response, request)
		return
	} else if path == "/upload" {
		UploadPic(response, request)
		return
	}

	http.NotFound(response, request) // 自己也是可以编写 404 这类的返回的 源码只是 func NotFound(w ResponseWriter, r *Request)
	return
}
