package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"fmt"
	"text/template"
	"bytes"
)

type DBConnectionInfo struct {
	Uname string
	DbName string
	Pwd string
	Host string
}

type Error struct {
	msg string
}

func (e *Error) Error() string {
	return fmt.Sprintf("has error -> %s", e.msg)
}

var sqlTemplate *template.Template

func init() {
	templatePoint, err := template.New("db").Parse("user={{.Uname}} dbname={{.DbName}} password={{.Pwd}} host={{.Host}} sslmode=disable")
	if nil != err {
		sqlTemplate = nil
	}
	sqlTemplate = templatePoint
}

func getConnection(info *DBConnectionInfo) (*sql.DB, error) {
	if nil == sqlTemplate {
		return nil, &Error{"初始化失败"}
	}

	var tpl bytes.Buffer
	if err := sqlTemplate.ExecuteTemplate(&tpl, "db", *info); err != nil {
		fmt.Println(err.Error())
		return nil, &Error{"sql转换失败"}
	}

	db, err := sql.Open("postgres", tpl.String())
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func main() {
	db, err := getConnection(&DBConnectionInfo{"centos", "administrative", "user123", "kini"})
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM administrative_area LIMIT $1", 10)
	if nil != err {
		fmt.Println(err.Error())
		return
	}
	for rows.Next() {
		var (
			id int
			level int
			basename string
			code string
			name string
			parent_id string
		)
		rows.Scan(&id, &level, &basename, &code, &name, &parent_id)
		fmt.Println(id, level, basename, code, name, parent_id)
	}

	db.Close()
}
