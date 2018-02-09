package main

import (
	"fmt"
	"log"
	"./DBUtil"
)

func main() {
	db, err := DBUtil.GetConnection(&DBUtil.DBConnectionInfo {
			"centos",
			"administrative",
			"user123",
			"kini"})
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
