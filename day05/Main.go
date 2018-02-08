package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"fmt"
)

func main() {
	connStr := "user=centos dbname=administrative password=user123 host=kini sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM administrative_area WHERE id = $1", 0)
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
}
