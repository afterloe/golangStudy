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

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

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

func Get(key string) interface{} {
	return packageJson[key]
}
