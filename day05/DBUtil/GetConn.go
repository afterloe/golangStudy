package DBUtil

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"fmt"
	"text/template"
	"bytes"
)

var sqlConnectionTemplate *template.Template

func init() {
	templatePoint, err := template.New("db").Parse("user={{.Uname}} dbname={{.DbName}} password={{.Pwd}} host={{.Host}} sslmode=disable")
	if nil != err {
		sqlConnectionTemplate = nil
	}
	sqlConnectionTemplate = templatePoint
}

func GetConnection(info *DBConnectionInfo) (*sql.DB, error) {
	if nil == sqlConnectionTemplate {
		return nil, &Error{"初始化失败"}
	}

	var tpl bytes.Buffer
	if err := sqlConnectionTemplate.ExecuteTemplate(&tpl, "db", *info); err != nil {
		fmt.Println(err.Error())
		return nil, &Error{"sql转换失败"}
	}

	db, err := sql.Open("postgres", tpl.String())
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}