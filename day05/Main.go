package main

import (
	"fmt"
	"./DBUtil"
)

func main() {
	DBUtil.InitConnectionInfo("centos", "administrative", "user123", "kini")
	mapResult, err := DBUtil.QueryInfo("SELECT * FROM administrative_area LIMIT $1", 10)
	if nil != err {
		fmt.Println(err.Error())
		return
	}
	for _, m := range mapResult {
		fmt.Println(m)
	}
}
