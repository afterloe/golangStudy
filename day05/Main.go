package main

import (
	"fmt"
	"./DBUtil"
)

func main() {
	// 初始化链接
	DBUtil.InitConnectionInfo("centos", "administrative", "user123", "kini")
	// 调用查询方法 -- 封装过 返回 map对象
	mapResult, err := DBUtil.QueryInfo("SELECT * FROM administrative_area LIMIT $1", 10)
	if nil != err {
		fmt.Println(err.Error())
		return
	}
	for _, m := range mapResult {
		fmt.Println(m)
	}
}
