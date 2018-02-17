package day07

import "testing"

func Test_inputInfo(t *testing.T) {
	str := "myName"
	if "afterloe" != str {
		t.Error("测试失败")
	} else {
		t.Log("测试通过")
	}
}

func Test_inputInfo_2(t *testing.T) {
	t.Error("这个还是不通过")
}