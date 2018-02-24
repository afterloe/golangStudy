package config

import (
	"../integrate/util"
	"../integrate/logger"
	"../exceptions"
	"os"
)

var packageJson map[string]interface{}

func checkError(err error) {
	if nil != err {
		logger.Error(err.Error())
		os.Exit(101)
		return
	}
}

func init() {
	configInfo, err := util.ReadRealFile("/Users/afterloe/Afterloe/go/golangStudy/day08", "package.json")
	checkError(err)
	pkg, err := util.FormatToStruct(&configInfo)
	checkError(err)
	if 0 == len(pkg) {
		checkError(&exceptions.Error{Msg: "read json fail", Code: 500})
	}
	packageJson = pkg
}

func Get(key string) interface{} {
	return packageJson[key]
}
