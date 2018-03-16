package config

import (
	"../integrate/util"
	"../integrate/logger"
	"../exceptions"
	"os"
	"path/filepath"
	"log"
	"strings"
)

var packageJson map[string]interface{}

func checkError(err error) {
	if nil != err {
		logger.Error(err.Error())
		os.Exit(101)
		return
	}
}

/**
 *  获取代码运行目录
 */
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

/**
 * 初始化包函数
 */
func init() {
	dir := getCurrentDirectory()
	configInfo, err := util.ReadRealFile(filepath.Join(dir, "./package.json"))
	checkError(err)
	pkg, err := util.FormatToStruct(&configInfo)
	checkError(err)
	if 0 == len(pkg) {
		checkError(&exceptions.Error{Msg: "read json fail", Code: 500})
	}
	packageJson = pkg
}

/**
 * 获取配置项
 *
 * @param key string 配置项key
 * @return interface{} 配置内容
 */
func Get(key string) interface{} {
	return packageJson[key]
}
