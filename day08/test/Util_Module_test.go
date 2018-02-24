package test

import (
	"testing"
	"../integrate/util"
)

func Test_UtilModule_readConfig(t *testing.T) {
	json, err := util.ReadRealFile("..", "package.json")
	if nil != err {
		t.Error(err.Error())
	}
	t.Log(json)
}

func Test_UtilModule_parseJSON(t *testing.T) {
	json, err := util.ReadRealFile("..", "package.json")
	if nil != err {
		t.Error(err.Error())
	}
	pkg, err := util.FormatToStruct(&json)
	if nil != err {
		t.Error(err.Error())
	}
	t.Log(pkg)
}