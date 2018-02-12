package DBUtil

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"fmt"
	"text/template"
	"bytes"
)

// 设置 sql 链接模板
var sqlConnectionTemplate *template.Template
var sqlConnectionStr string

func init() {
	// 初始化模板参数
	templatePoint, _ := template.New("db").Parse("user={{.Uname}} dbname={{.DbName}} password={{.Pwd}} host={{.Host}} sslmode=disable")
	sqlConnectionTemplate = templatePoint
}

type breakthroughPoint interface {
	execute(*sql.DB) (interface{}, error)
}

func UseConnection(point *breakthroughPoint) (interface{}, error) {
	db, error := openConnection()
	if nil != error {
		return nil, &Error{error.Error()}
	}
	point.execute(db)

	db.Close()

	return nil, nil
}

func openConnection()(*sql.DB, error) {
	if "" != sqlConnectionStr {
		if nil == sqlConnectionTemplate {
			return nil, &Error{"初始化失败"}
		}
		connectInfo := getConnectionInfo()
		var tpl bytes.Buffer
		if err := sqlConnectionTemplate.ExecuteTemplate(&tpl, "db", connectInfo); err != nil {
			fmt.Println(err.Error())
			return nil, &Error{"sql转换失败"}
		}

		sqlConnectionStr = tpl.String()
	}

	return sql.Open("postgres", sqlConnectionStr)
}