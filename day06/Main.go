package main

import (
	"fmt"
	"net/http"
	"log"
	"./simpleService"
)

func main() {
	serverListenAddr := simpleService.POINT

	// 设置 访问路由
	/*
	非 interface 的函数绑定，也许会更合适day05 中的 DBUtil 封装的方法

		func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
			DefaultServeMux.HandleFunc(pattern, handler)
		}
	 */
	http.HandleFunc("/", simpleService.Execute)
	// 开始监听服务
	err := http.ListenAndServe(serverListenAddr, nil)
	if nil != err {
		log.Fatal("Listen Port is", serverListenAddr, "Failed! because", err)
	}
	fmt.Println("service is success running ... in " + serverListenAddr)
}
