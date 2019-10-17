package DBUtil

import "fmt"

// 定义链接信息结构体
type connectInfo struct {
	Uname string
	DbName string
	Pwd string
	Host string
}

// 饿汉 模式
var info interface{}

// 重写 toString 方法，让他能够输出 描写的信息
func (info *connectInfo) String() string {
	return fmt.Sprintf("{Uname: %s, DbName: %s, Pwd: %s, Host: %s}",
		info.Uname, info.DbName, info.Pwd, info.Host)
}

// 初始化链接参数 如果已经初始化 则可以重新设置这个参数
func InitConnectionInfo(uname, dbname, pwd, host string) {
	if val, ok := info.(*connectInfo); !ok {
		info = &connectInfo{uname, dbname, pwd, host}
		return
	} else {
		val.Uname = uname
		val.DbName = dbname
		val.Pwd = pwd
		val.Host = host
	}
}

// 输出全局设置的 链接信息
func getConnectionInfo() *connectInfo {
	if val, ok := info.(*connectInfo); ok {
		return val
	} else {
		return nil
	}
}