package day07

import "testing"

// 普通测试
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

// 压力测试
func Benchmark_Input(t *testing.B) {
	for i:=0; i < t.N; i++ {
		// do some thing
	}
}

func Benchmark_Input2(t *testing.B) {
	t.StopTimer() // 暂停 计数器
	// 初始化资源
	t.StartTimer() // 启动计时器
	for i:=0; i < t.N; i++ {
		// do some thing
	}
}