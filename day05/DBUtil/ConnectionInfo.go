package DBUtil

import (
	"fmt"
)

type connectInfo struct {
	Uname string
	DbName string
	Pwd string
	Host string
}

var info interface{}

func (info *connectInfo) String() string {
	return fmt.Sprintf("{Uname: %s, DbName: %s, Pwd: %s, Host: %s}",
		info.Uname, info.DbName, info.Pwd, info.Host)
}

func InitConnectionInfo(uname, dbname, pwd, host string) {
	info = &connectInfo{uname, dbname, pwd, host}
}

func getConnectionInfo() *connectInfo {
	if val, ok := info.(*connectInfo); ok {
		return val
	} else {
		return nil
	}
}