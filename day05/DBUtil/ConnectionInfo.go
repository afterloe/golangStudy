package DBUtil

import "fmt"

type DBConnectionInfo struct {
	Uname string
	DbName string
	Pwd string
	Host string
}

func (info *DBConnectionInfo) String() string {
	return fmt.Sprintf("{Uname: %s, DbName: %s, Pwd: %s, Host: %s}",
		info.Uname, info.DbName, info.Pwd, info.Host)
}