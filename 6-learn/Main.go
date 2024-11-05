package main

import (
	"fmt"
	"net/http"
	"log"
	//"./simpleService"
	"./simpleService2"
)

/*
非 interface 的函数绑定，也许会更合适day05 中的 DBUtil 封装的方法

	func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
		DefaultServeMux.HandleFunc(pattern, handler)
	}
 */
func main() {
	//serverListenAddr := simpleService.POINT // for simpleService - 1
	cynomys := &simpleService2.CustomHandler{ServerListenAddr: "127.0.0.1:8081"} // for simpleService - 2 自己定义路由表
	serverListenAddr := cynomys.ServerListenAddr
	// 设置 访问路由
	//http.HandleFunc("/", simpleService.Execute) // for simpleService - 1

	// 开始监听服务
	//err := http.ListenAndServe(serverListenAddr, nil) // for simpleService - 1 默认的路由表是没有404的
	err := http.ListenAndServe(serverListenAddr, cynomys) // for simpleService - 2 这里使用了 自己定义的路由表
	fmt.Println("service is success running ... in " + serverListenAddr)
	if nil != err {
		log.Fatal("Listen Port is", serverListenAddr, "Failed! because", err)
	}
}
