# 第六条

go中是可以指定func传入的方式来控制输入输出的，例如如下代码
```golang
	func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
		DefaultServeMux.HandleFunc(pattern, handler)
	}
```
可以不是用interface就可以实现特定函数的传入，在很多地方是可以这么实用的，具体和interface的区别，我还没发现就是了。  

go使用http包来设置http服务，该服务可以使用自己的默认路由，启动后采用`HandleFunc`方法来设置每条路由的处理函数，默认情况下go是不会自己自动解析request中的参数的，需要手动调用`request.ParseForm()` 方法来实现解析。该方法只需要调度一次即可，多次调用是不会有什么结果的。然后`request.URL.Path`返回的是访问的相对路径，而`request.Form["name"]`可以获取输入参数为name的值，返回的是一个[]interface{}数组。所以使用的时候需要注意一下。具体代码可以参照[这里](./simpleService/SimpleService.go)  

这是启动服务之后进行的postman的测试脚本如下
```sbtshell
curl --request POST \
  --url 'http://127.0.0.1:8082/?name=afterloe' \
  --header 'Cache-Control: no-cache' \
  --header 'Content-Type: application/x-www-form-urlencoded' \
  --header 'Postman-Token: 57315e28-b94f-df3f-721c-127450eac187' \
  --data 'name=joe&age=12'
```

```sbtshell
curl --request POST \
  --url 'http://127.0.0.1:8081/v1/user?name=afterloe' \
  --header 'Cache-Control: no-cache' \
  --header 'Content-Type: application/x-www-form-urlencoded' \
  --header 'Postman-Token: 9dc41a2e-6f9a-9f92-d3e8-a316055f9956' \
  --data 'name=joe&age=12'
```

simpleService2里面不再使用默认的路由解析表了，这里采用的是自己定义的路由解析类

```golang
type CustomHandler struct {
	ServerListenAddr string
}
```

该结构体需要实现一个接口

```golang
func (*CustomHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {}
```

在这个函数里面写入响应的路径，默认可以调用`http.NotFound(response, request)` 来输出返回，这个是可以自己定义的。  

文件上传也很容易，可以使用`file, handler, error := request.FormFile("icon")` 来进行文件的转换。在使用之前需要设置缓冲区大小`request.ParseMultipartForm()` 单位是字节。然后再用`io.Copy`方法就可以了。具体代码可以参看[这里](./simpleService2/SimpleService.go)  

这下面针对service2里面的服务一个postman的测试脚本

```sbtshell
curl --request POST \
  --url http://127.0.0.1:8081/ \
  --header 'Cache-Control: no-cache' \
  --header 'Content-Type: application/x-www-form-urlencoded' \
  --header 'Postman-Token: 3c45ca83-d791-c3ba-9144-33dc87070627' \
  --data age=12
```

```sbtshell
curl --request POST \
  --url http://127.0.0.1:8081/ \
  --header 'Cache-Control: no-cache' \
  --header 'Content-Type: application/x-www-form-urlencoded' \
  --header 'Postman-Token: 30a9abe9-a19c-db3a-484d-3b7efe6d55bf' \
  --data 'name=joe&age=12'
```

最后是安全问题  
 这里推荐scrypt方案，scrypt是由著名的FreeBSD黑客Colin Percival为他的备份服务Tarsnap开发的。 [库](http://code.google.com/p/go/source/browse?repo=crypto#hg%2Fscrypt)
